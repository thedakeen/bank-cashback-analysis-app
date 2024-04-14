package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *CardModel) SetCard(userId primitive.ObjectID, card_number, card_type, bank_name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insert := bson.M{
		"card_number": card_number,
		"card_type":   card_type,
		"bank_name":   bank_name,
		"user_id":     userId,
	}

	_, err := m.C.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$push": bson.M{"cards": insert}})
	if err != nil {
		return err
	}

	return nil
}
