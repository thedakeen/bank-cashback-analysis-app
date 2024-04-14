package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CardModel struct {
	C *mongo.Collection
}

func NewCardModel(cardCollection *mongo.Collection) *CardModel {
	return &CardModel{
		C: cardCollection,
	}
}
