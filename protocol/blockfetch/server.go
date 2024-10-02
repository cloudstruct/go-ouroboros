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

package blockfetch

import (
	"fmt"

	"github.com/blinklabs-io/gouroboros/cbor"
	"github.com/blinklabs-io/gouroboros/protocol"
)

type Server struct {
	*protocol.Protocol
	config          *Config
	callbackContext CallbackContext
	protoOptions    protocol.ProtocolOptions
}

func NewServer(protoOptions protocol.ProtocolOptions, cfg *Config) *Server {
	s := &Server{
		config: cfg,
		// Save this for re-use later
		protoOptions: protoOptions,
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
		InitialState:        StateIdle,
	}
	s.Protocol = protocol.New(protoConfig)
}

func (s *Server) NoBlocks() error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("server called %s NoBlocks()", ProtocolName))
	msg := NewMsgNoBlocks()
	return s.SendMessage(msg)
}

func (s *Server) StartBatch() error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("server called %s StartBatch()", ProtocolName))
	msg := NewMsgStartBatch()
	return s.SendMessage(msg)
}

func (s *Server) Block(blockType uint, blockData []byte) error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("server called %s Block(blockType: %+v, blockData: %x)", ProtocolName, blockType, blockData))
	wrappedBlock := WrappedBlock{
		Type:     blockType,
		RawBlock: blockData,
	}
	wrappedBlockData, err := cbor.Encode(&wrappedBlock)
	if err != nil {
		return err
	}
	msg := NewMsgBlock(wrappedBlockData)
	return s.SendMessage(msg)
}

func (s *Server) BatchDone() error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("server called %s BatchDone()", ProtocolName))
	msg := NewMsgBatchDone()
	return s.SendMessage(msg)
}

func (s *Server) messageHandler(msg protocol.Message) error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("handling server message for %s", ProtocolName))
	var err error
	switch msg.Type() {
	case MessageTypeRequestRange:
		err = s.handleRequestRange(msg)
	case MessageTypeClientDone:
		err = s.handleClientDone()
	default:
		err = fmt.Errorf(
			"%s: received unexpected message type %d",
			ProtocolName,
			msg.Type(),
		)
	}
	return err
}

func (s *Server) handleRequestRange(msg protocol.Message) error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("handling server request range for %s", ProtocolName))
	if s.config == nil || s.config.RequestRangeFunc == nil {
		return fmt.Errorf(
			"received block-fetch RequestRange message but no callback function is defined",
		)
	}
	msgRequestRange := msg.(*MsgRequestRange)
	return s.config.RequestRangeFunc(
		s.callbackContext,
		msgRequestRange.Start,
		msgRequestRange.End,
	)
}

func (s *Server) handleClientDone() error {
	s.Protocol.Logger().
		Debug(fmt.Sprintf("handling server client done for %s", ProtocolName))
	// Restart protocol
	s.Protocol.Stop()
	s.initProtocol()
	s.Protocol.Start()
	return nil
}
