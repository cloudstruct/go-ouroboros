// Copyright 2024 Blink Labs Software
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package txsubmission

import (
	"fmt"
	"sync"

	"github.com/blinklabs-io/gouroboros/protocol"
)

// Server implements the TxSubmission server
type Server struct {
	*protocol.Protocol
	config                 *Config
	callbackContext        CallbackContext
	protoOptions           protocol.ProtocolOptions
	ackCount               int
	requestTxIdsResultChan chan []TxIdAndSize
	requestTxsResultChan   chan []TxBody
	onceStart              sync.Once
}

// NewServer returns a new TxSubmission server object
func NewServer(protoOptions protocol.ProtocolOptions, cfg *Config) *Server {
	s := &Server{
		config: cfg,
		// Save this for re-use later
		protoOptions:           protoOptions,
		requestTxIdsResultChan: make(chan []TxIdAndSize),
		requestTxsResultChan:   make(chan []TxBody),
	}
	s.callbackContext = CallbackContext{
		Server:       s,
		ConnectionId: protoOptions.ConnectionId,
	}
	s.initProtocol()
	return s
}

func (s *Server) initProtocol() {
	protoConfig := protocol.ProtocolConfig{
		Name:                ProtocolName,
		ProtocolId:          ProtocolId,
		Muxer:               s.protoOptions.Muxer,
		Logger:              s.protoOptions.Logger,
		ErrorChan:           s.protoOptions.ErrorChan,
		Mode:                s.protoOptions.Mode,
		Role:                protocol.ProtocolRoleServer,
		MessageHandlerFunc:  s.messageHandler,
		MessageFromCborFunc: NewMsgFromCbor,
		StateMap:            StateMap,
		InitialState:        stateInit,
	}
	s.Protocol = protocol.New(protoConfig)
}

func (s *Server) Start() {
	s.onceStart.Do(func() {
		s.Protocol.Logger().
			Debug("starting server protocol",
				"component", "network",
				"protocol", ProtocolName,
				"connection_id", s.callbackContext.ConnectionId.String(),
			)
		s.Protocol.Start()
		// Start goroutine to cleanup resources on protocol shutdown
		go func() {
			<-s.Protocol.DoneChan()
			close(s.requestTxIdsResultChan)
			close(s.requestTxsResultChan)
		}()
	})
}

// RequestTxIds requests the next set of TX identifiers from the remote node's mempool
func (s *Server) RequestTxIds(
	blocking bool,
	reqCount int,
) ([]TxIdAndSize, error) {
	s.Protocol.Logger().
		Debug(
			fmt.Sprintf("calling RequestTxIds(blocking: %+v, reqCount: %d)", blocking, reqCount),
			"component", "network",
			"protocol", ProtocolName,
			"role", "server",
			"connection_id", s.callbackContext.ConnectionId.String(),
		)
	msg := NewMsgRequestTxIds(blocking, uint16(s.ackCount), uint16(reqCount))
	if err := s.SendMessage(msg); err != nil {
		return nil, err
	}
	// Wait for result
	txIds, ok := <-s.requestTxIdsResultChan
	if !ok {
		return nil, protocol.ProtocolShuttingDownError
	}
	// Update ack count for next call
	s.ackCount = len(txIds)
	return txIds, nil
}

// RequestTxs requests the content of the requested TX identifiers from the remote node's mempool
func (s *Server) RequestTxs(txIds []TxId) ([]TxBody, error) {
	s.Protocol.Logger().
		Debug(
			fmt.Sprintf("calling RequestTxs(txIds: %+v)", txIds),
			"component", "network",
			"protocol", ProtocolName,
			"role", "server",
			"connection_id", s.callbackContext.ConnectionId.String(),
		)
	msg := NewMsgRequestTxs(txIds)
	if err := s.SendMessage(msg); err != nil {
		return nil, err
	}
	// Wait for result
	txs, ok := <-s.requestTxsResultChan
	if !ok {
		return nil, protocol.ProtocolShuttingDownError
	}
	return txs, nil
}

func (s *Server) messageHandler(msg protocol.Message) error {
	var err error
	switch msg.Type() {
	case MessageTypeReplyTxIds:
		err = s.handleReplyTxIds(msg)
	case MessageTypeReplyTxs:
		err = s.handleReplyTxs(msg)
	case MessageTypeDone:
		err = s.handleDone()
	case MessageTypeInit:
		err = s.handleInit()
	default:
		err = fmt.Errorf(
			"%s: received unexpected message type %d",
			ProtocolName,
			msg.Type(),
		)
	}
	return err
}

func (s *Server) handleReplyTxIds(msg protocol.Message) error {
	s.Protocol.Logger().
		Debug("reply tx ids",
			"component", "network",
			"protocol", ProtocolName,
			"role", "server",
			"connection_id", s.callbackContext.ConnectionId.String(),
		)
	// Check for shutdown
	select {
	case <-s.Protocol.DoneChan():
		return protocol.ProtocolShuttingDownError
	default:
	}
	msgReplyTxIds := msg.(*MsgReplyTxIds)
	s.requestTxIdsResultChan <- msgReplyTxIds.TxIds
	return nil
}

func (s *Server) handleReplyTxs(msg protocol.Message) error {
	s.Protocol.Logger().
		Debug("reply txs",
			"component", "network",
			"protocol", ProtocolName,
			"role", "server",
			"connection_id", s.callbackContext.ConnectionId.String(),
		)
	// Check for shutdown
	select {
	case <-s.Protocol.DoneChan():
		return protocol.ProtocolShuttingDownError
	default:
	}
	msgReplyTxs := msg.(*MsgReplyTxs)
	s.requestTxsResultChan <- msgReplyTxs.Txs
	return nil
}

func (s *Server) handleDone() error {
	s.Protocol.Logger().
		Debug("done",
			"component", "network",
			"protocol", ProtocolName,
			"role", "server",
			"connection_id", s.callbackContext.ConnectionId.String(),
		)
	// Restart protocol
	s.Protocol.Stop()
	s.initProtocol()
	s.requestTxIdsResultChan = make(chan []TxIdAndSize)
	s.requestTxsResultChan = make(chan []TxBody)
	s.Protocol.Start()
	return nil
}

func (s *Server) handleInit() error {
	s.Protocol.Logger().
		Debug("init",
			"component", "network",
			"protocol", ProtocolName,
			"role", "server",
			"connection_id", s.callbackContext.ConnectionId.String(),
		)
	if s.config == nil || s.config.InitFunc == nil {
		return fmt.Errorf(
			"received tx-submission Init message but no callback function is defined",
		)
	}
	// Call the user callback function
	return s.config.InitFunc(s.callbackContext)
}
