package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TransactionHash string             `bson:"transaction_hash" json:"transaction_hash"`
	Time            time.Time          `bson:"time" json:"time"`
	Fee             int64              `bson:"fee" json:"fee"`
	Confirmations   *int64             `json:"confirmations" bson:"confirmations"`
	BlockHeight     *int64             `json:"block_height" bson:"block_height"`
	BlockIndex      *int64             `json:"block_index" bson:"block_index"`
	Inputs          []Inputs           `bson:"inputs" json:"inputs"`
	Outputs         []Outputs          `bson:"outputs" json:"outputs"`
}

type Inputs struct {
	PreviousTxid int64  `bson:"previous_txid" json:"previous_txid"`
	Index        int64  `bson:"index" json:"index"`
	Address      string `bson:"address" json:"address"`
	Value        int64  `bson:"value" json:"value"`
}

type Outputs struct {
	Address string `bson:"address" json:"address"`
	Value   int64  `bson:"value" json:"value"`
}
