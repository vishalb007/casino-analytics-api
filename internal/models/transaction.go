package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedAt time.Time            `bson:"createdAt"`
	UserID    primitive.ObjectID   `bson:"userId"`
	RoundID   string               `bson:"roundId"`
	Type      string               `bson:"type"`
	Amount    primitive.Decimal128 `bson:"amount"`
	Currency  string               `bson:"currency"`
	USDAmount primitive.Decimal128 `bson:"usdAmount"`
}