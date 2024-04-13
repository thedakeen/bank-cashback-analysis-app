package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CardModel struct {
	C *mongo.Collection
}

func NewCardModel(cardCollection *mongo.Collection) *CardModel {
	return &CardModel{
		C: cardCollection,
	}
}

func (m *CardModel) SetCard(card_number, card_type, bank_name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insert := bson.M{
		"card_number": card_number,
		"card_type":   card_type,
		"bank_name":   bank_name,
	}

	_, err := m.C.InsertOne(ctx, insert)
	if err != nil {
		return err
	}

	return nil
}
