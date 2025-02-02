// Copyright 2025 Blink Labs Software
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

package conway_test

import (
	"encoding/hex"
	"github.com/blinklabs-io/gouroboros/ledger/common"
	"github.com/blinklabs-io/gouroboros/ledger/conway"
	"testing"
)

// Placeholder for sample CBOR data
var sampleCborData = []byte{ /* ... */ }

func TestConwayBlock_UnmarshalCBOR(t *testing.T) {
	var block conway.ConwayBlock
	err := block.UnmarshalCBOR(sampleCborData)
	if err != nil {
		t.Fatalf("UnmarshalCBOR failed: %v", err)
	}
}

func TestConwayBlock_Hash(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	hash := block.Hash()
	expectedHash := "expected hash value"
	if hash != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, hash)
	}
}

func TestConwayBlock_Header(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	header := block.Header()
	if header == nil {
		t.Error("expected non-nil header")
	}
}

func TestConwayBlock_PrevHash(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	prevHash := block.PrevHash()
	expectedPrevHash := "expected prev hash value"
	if prevHash != expectedPrevHash {
		t.Errorf("expected %s, got %s", expectedPrevHash, prevHash)
	}
}

func TestConwayBlock_BlockNumber(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	blockNumber := block.BlockNumber()
	expectedBlockNumber := uint64(123456) // example value
	if blockNumber != expectedBlockNumber {
		t.Errorf("expected %d, got %d", expectedBlockNumber, blockNumber)
	}
}

func TestConwayBlock_SlotNumber(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	slotNumber := block.SlotNumber()
	expectedSlotNumber := uint64(654321) // example value
	if slotNumber != expectedSlotNumber {
		t.Errorf("expected %d, got %d", expectedSlotNumber, slotNumber)
	}
}

func TestConwayBlock_IssuerVkey(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	issuerVkey := block.IssuerVkey()
	expectedIssuerVkey := common.IssuerVkey{}
	if issuerVkey != expectedIssuerVkey {
		t.Errorf("expected %v, got %v", expectedIssuerVkey, issuerVkey)
	}
}

func TestConwayBlock_BlockBodySize(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	blockBodySize := block.BlockBodySize()
	expectedBlockBodySize := uint64(1024) // example value
	if blockBodySize != expectedBlockBodySize {
		t.Errorf("expected %d, got %d", expectedBlockBodySize, blockBodySize)
	}
}

func TestConwayBlock_Era(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	era := block.Era()
	expectedEra := conway.EraConway
	if era != expectedEra {
		t.Errorf("expected %v, got %v", expectedEra, era)
	}
}

func TestConwayBlock_Transactions(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	transactions := block.Transactions()
	if len(transactions) != len(block.TransactionBodies) {
		t.Errorf("expected %d transactions, got %d", len(block.TransactionBodies), len(transactions))
	}
}

func TestConwayBlock_Utxorpc(t *testing.T) {
	var block conway.ConwayBlock
	block.UnmarshalCBOR(sampleCborData)
	utxorpcBlock := block.Utxorpc()
	if utxorpcBlock == nil {
		t.Error("expected non-nil utxorpc.Block")
	}
}

func TestConwayBlockHeader_Era(t *testing.T) {
	var header conway.ConwayBlockHeader
	era := header.Era()
	expectedEra := conway.EraConway
	if era != expectedEra {
		t.Errorf("expected %v, got %v", expectedEra, era)
	}
}

func TestConwayRedeemers_UnmarshalCBOR(t *testing.T) {
	var redeemers conway.ConwayRedeemers
	err := redeemers.UnmarshalCBOR(sampleCborData)
	if err != nil {
		t.Fatalf("UnmarshalCBOR failed: %v", err)
	}
}

func TestConwayTransactionWitnessSet_UnmarshalCBOR(t *testing.T) {
	var witnessSet conway.ConwayTransactionWitnessSet
	err := witnessSet.UnmarshalCBOR(sampleCborData)
	if err != nil {
		t.Fatalf("UnmarshalCBOR failed: %v", err)
	}
}

func TestConwayTransactionInputSet_UnmarshalCBOR(t *testing.T) {
	var inputSet conway.ConwayTransactionInputSet
	err := inputSet.UnmarshalCBOR(sampleCborData)
	if err != nil {
		t.Fatalf("UnmarshalCBOR failed: %v", err)
	}
}

func TestConwayTransactionBody_UnmarshalCBOR(t *testing.T) {
	var body conway.ConwayTransactionBody
	err := body.UnmarshalCBOR(sampleCborData)
	if err != nil {
		t.Fatalf("UnmarshalCBOR failed: %v", err)
	}
}

func TestConwayTransactionBody_Inputs(t *testing.T) {
	var body conway.ConwayTransactionBody
	body.UnmarshalCBOR(sampleCborData)
	inputs := body.Inputs()
	if len(inputs) != len(body.TxInputs.Items()) {
		t.Errorf("expected %d inputs, got %d", len(body.TxInputs.Items()), len(inputs))
	}
}

func TestConwayTransactionBody_ProtocolParameterUpdates(t *testing.T) {
	var body conway.ConwayTransactionBody
	body.UnmarshalCBOR(sampleCborData)
	epoch, updates := body.ProtocolParameterUpdates()
	if epoch != body.Update.Epoch {
		t.Errorf("expected epoch %d, got %d", body.Update.Epoch, epoch)
	}
	if len(updates) != len(body.Update.ProtocolParamUpdates) {
		t.Errorf("expected %d updates, got %d", len(body.Update.ProtocolParamUpdates), len(updates))
	}
}

func TestConwayTransactionBody_VotingProcedures(t *testing.T) {
	var body conway.ConwayTransactionBody
	body.UnmarshalCBOR(sampleCborData)
	votingProcedures := body.VotingProcedures()
	if votingProcedures != body.TxVotingProcedures {
		t.Errorf("expected %v, got %v", body.TxVotingProcedures, votingProcedures)
	}
}

func TestConwayTransactionBody_ProposalProcedures(t *testing.T) {
	var body conway.ConwayTransactionBody
	body.UnmarshalCBOR(sampleCborData)
	proposalProcedures := body.ProposalProcedures()
	if len(proposalProcedures) != len(body.TxProposalProcedures) {
		t.Errorf("expected %d proposal procedures, got %d", len(body.TxProposalProcedures), len(proposalProcedures))
	}
}

func TestConwayTransactionBody_CurrentTreasuryValue(t *testing.T) {
	var body conway.ConwayTransactionBody
	body.UnmarshalCBOR(sampleCborData)
	currentTreasuryValue := body.CurrentTreasuryValue()
	if currentTreasuryValue != body.TxCurrentTreasuryValue {
		t.Errorf("expected %d, got %d", body.TxCurrentTreasuryValue, currentTreasuryValue)
	}
}

func TestConwayTransactionBody_Donation(t *testing.T) {
	var body conway.ConwayTransactionBody
	body.UnmarshalCBOR(sampleCborData)
	donation := body.Donation()
	if donation != body.TxDonation {
		t.Errorf("expected %d, got %d", body.TxDonation, donation)
	}
}

func TestConwayTransaction_Type(t *testing.T) {
	tx := conway.ConwayTransaction{}
	txType := tx.Type()
	expectedTxType := conway.TxTypeConway
	if txType != expectedTxType {
		t.Errorf("expected %d, got %d", expectedTxType, txType)
	}
}

func TestConwayTransaction_Hash(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	hash := tx.Hash()
	expectedHash := "expected hash value"
	if hash != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, hash)
	}
}

func TestConwayTransaction_Inputs(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	inputs := tx.Inputs()
	if len(inputs) != len(tx.Body.TxInputs.Items()) {
		t.Errorf("expected %d inputs, got %d", len(tx.Body.TxInputs.Items()), len(inputs))
	}
}

func TestConwayTransaction_Outputs(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	outputs := tx.Outputs()
	if len(outputs) != len(tx.Body.TxOutputs) {
		t.Errorf("expected %d outputs, got %d", len(tx.Body.TxOutputs), len(outputs))
	}
}

func TestConwayTransaction_Fee(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	fee := tx.Fee()
	expectedFee := uint64(1000) // example value
	if fee != expectedFee {
		t.Errorf("expected %d, got %d", expectedFee, fee)
	}
}

func TestConwayTransaction_TTL(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	ttl := tx.TTL()
	expectedTTL := uint64(10000) // example value
	if ttl != expectedTTL {
		t.Errorf("expected %d, got %d", expectedTTL, ttl)
	}
}

func TestConwayTransaction_ValidityIntervalStart(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	validityIntervalStart := tx.ValidityIntervalStart()
	expectedValidityIntervalStart := uint64(5000) // example value
	if validityIntervalStart != expectedValidityIntervalStart {
		t.Errorf("expected %d, got %d", expectedValidityIntervalStart, validityIntervalStart)
	}
}

func TestConwayTransaction_ProtocolParameterUpdates(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	epoch, updates := tx.ProtocolParameterUpdates()
	if epoch != tx.Body.Update.Epoch {
		t.Errorf("expected epoch %d, got %d", tx.Body.Update.Epoch, epoch)
	}
	if len(updates) != len(tx.Body.Update.ProtocolParamUpdates) {
		t.Errorf("expected %d updates, got %d", len(tx.Body.Update.ProtocolParamUpdates), len(updates))
	}
}

func TestConwayTransaction_ReferenceInputs(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	referenceInputs := tx.ReferenceInputs()
	if len(referenceInputs) != len(tx.Body.ReferenceInputs()) {
		t.Errorf("expected %d reference inputs, got %d", len(tx.Body.ReferenceInputs()), len(referenceInputs))
	}
}

func TestConwayTransaction_Collateral(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	collateral := tx.Collateral()
	if len(collateral) != len(tx.Body.Collateral()) {
		t.Errorf("expected %d collateral inputs, got %d", len(tx.Body.Collateral()), len(collateral))
	}
}

func TestConwayTransaction_CollateralReturn(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	collateralReturn := tx.CollateralReturn()
	if collateralReturn != tx.Body.CollateralReturn() {
		t.Errorf("expected %v, got %v", tx.Body.CollateralReturn(), collateralReturn)
	}
}

func TestConwayTransaction_TotalCollateral(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	totalCollateral := tx.TotalCollateral()
	if totalCollateral != tx.Body.TotalCollateral() {
		t.Errorf("expected %d, got %d", tx.Body.TotalCollateral(), totalCollateral)
	}
}

func TestConwayTransaction_Certificates(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	certificates := tx.Certificates()
	if len(certificates) != len(tx.Body.Certificates()) {
		t.Errorf("expected %d certificates, got %d", len(tx.Body.Certificates()), len(certificates))
	}
}

func TestConwayTransaction_Withdrawals(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	withdrawals := tx.Withdrawals()
	if len(withdrawals) != len(tx.Body.Withdrawals()) {
		t.Errorf("expected %d withdrawals, got %d", len(tx.Body.Withdrawals()), len(withdrawals))
	}
}

func TestConwayTransaction_AuxDataHash(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	auxDataHash := tx.AuxDataHash()
	if auxDataHash != tx.Body.AuxDataHash() {
		t.Errorf("expected %v, got %v", tx.Body.AuxDataHash(), auxDataHash)
	}
}

func TestConwayTransaction_RequiredSigners(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	requiredSigners := tx.RequiredSigners()
	if len(requiredSigners) != len(tx.Body.RequiredSigners()) {
		t.Errorf("expected %d required signers, got %d", len(tx.Body.RequiredSigners()), len(requiredSigners))
	}
}

func TestConwayTransaction_AssetMint(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	assetMint := tx.AssetMint()
	if assetMint != tx.Body.AssetMint() {
		t.Errorf("expected %v, got %v", tx.Body.AssetMint(), assetMint)
	}
}

func TestConwayTransaction_ScriptDataHash(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	scriptDataHash := tx.ScriptDataHash()
	if scriptDataHash != tx.Body.ScriptDataHash() {
		t.Errorf("expected %v, got %v", tx.Body.ScriptDataHash(), scriptDataHash)
	}
}

func TestConwayTransaction_VotingProcedures(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	votingProcedures := tx.VotingProcedures()
	if votingProcedures != tx.Body.VotingProcedures() {
		t.Errorf("expected %v, got %v", tx.Body.VotingProcedures(), votingProcedures)
	}
}

func TestConwayTransaction_ProposalProcedures(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	proposalProcedures := tx.ProposalProcedures()
	if len(proposalProcedures) != len(tx.Body.ProposalProcedures()) {
		t.Errorf("expected %d proposal procedures, got %d", len(tx.Body.ProposalProcedures()), len(proposalProcedures))
	}
}

func TestConwayTransaction_CurrentTreasuryValue(t *testing.T) {
	var tx conway.ConwayTransaction
	tx.UnmarshalCBOR(sampleCborData)
	currentTreasuryValue := tx.CurrentTreasuryValue()
	if currentTreasuryValue != tx.Body.CurrentTreasuryValue() {
		t.Errorf("expected %d, got %d", tx.Body.CurrentTreasuryValue(), currentTreasuryValue)
	}
}
