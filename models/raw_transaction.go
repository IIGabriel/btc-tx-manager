package models

import (
	"time"
)

type RawTransaction struct {
	Hash        string     `bson:"hash" json:"hash"`
	Ver         int        `bson:"ver" json:"ver"`
	VinSz       int        `bson:"vinSz" json:"vin_sz"`
	VoutSz      int        `bson:"voutSz" json:"vout_sz"`
	Size        int        `bson:"size" json:"size"`
	Weight      int        `bson:"weight" json:"weight"`
	Fee         int        `bson:"fee" json:"fee"`
	RelayedBy   string     `bson:"relayedBy" json:"relayed_by"`
	LockTime    int        `bson:"lockTime" json:"lock_time"`
	TxIndex     int64      `bson:"txIndex" json:"tx_index"`
	DoubleSpend bool       `bson:"doubleSpend" json:"double_spend"`
	Time        int64      `bson:"time" json:"time"`
	BlockIndex  *int64     `bson:"blockIndex" json:"block_index"`
	BlockHeight *int64     `bson:"blockHeight" json:"block_height"`
	Inputs      []RawInput `bson:"inputs" json:"inputs"`
	Out         []RawOut   `bson:"out" json:"out"`
}

type RawInput struct {
	Sequence int64   `bson:"sequence" json:"sequence"`
	Witness  string  `bson:"witness" json:"witness"`
	Script   string  `bson:"script" json:"script"`
	Index    int     `bson:"index" json:"index"`
	PrevOut  PrevOut `bson:"prevOut" json:"prev_out"`
}

type RawOut struct {
	Type              int                 `bson:"type" json:"type"`
	Spent             bool                `bson:"spent" json:"spent"`
	Value             int64               `bson:"value" json:"value"`
	SpendingOutpoints []SpendingOutpoints `bson:"spendingOutpoints" json:"spending_outpoints"`
	N                 int                 `bson:"n" json:"n"`
	TxIndex           int64               `bson:"txIndex" json:"tx_index"`
	Script            string              `bson:"script" json:"script"`
	Addr              string              `bson:"addr" json:"addr"`
}

type PrevOut struct {
	Addr              string              `bson:"addr" json:"addr"`
	N                 int                 `bson:"n" json:"n"`
	Script            string              `bson:"script" json:"script"`
	SpendingOutpoints []SpendingOutpoints `bson:"spendingOutpoints" json:"spending_outpoints"`
	Spent             bool                `bson:"spent" json:"spent"`
	TxIndex           int64               `bson:"txIndex" json:"tx_index"`
	Type              int                 `bson:"type" json:"type"`
	Value             int64               `bson:"value" json:"value"`
}

type SpendingOutpoints struct {
	N       int   `bson:"n" json:"n"`
	TxIndex int64 `bson:"txIndex" json:"tx_index"`
}

func (rawTx RawTransaction) ConvertToTransaction() *Transaction {
	var inputs []Inputs
	for _, rawInput := range rawTx.Inputs {
		input := Inputs{
			PreviousTxid: rawInput.PrevOut.TxIndex,
			Index:        int64(rawInput.Index),
			Address:      rawInput.PrevOut.Addr,
			Value:        rawInput.PrevOut.Value,
		}
		inputs = append(inputs, input)
	}

	var outputs []Outputs
	for _, rawOutput := range rawTx.Out {
		output := Outputs{
			Address: rawOutput.Addr,
			Value:   rawOutput.Value,
		}
		outputs = append(outputs, output)
	}

	tx := Transaction{
		TransactionHash: rawTx.Hash,
		Time:            time.Unix(rawTx.Time, 0),
		Fee:             int64(rawTx.Fee),
		BlockHeight:     rawTx.BlockHeight,
		BlockIndex:      rawTx.BlockIndex,
		Inputs:          inputs,
		Outputs:         outputs,
	}

	return &tx
}
