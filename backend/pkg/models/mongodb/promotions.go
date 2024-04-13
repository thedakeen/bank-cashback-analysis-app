package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PromoModel struct {
	C *mongo.Collection
}

func NewPromoModel(c *mongo.Collection) *PromoModel {
	return &PromoModel{C: c}
}

func (m *PromoModel) AddKaspi(title, source_url, bank_name, promo_type, category string, bonus_rate float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insert := bson.M{
		"title":      title,
		"source_url": source_url,
		"bank_name":  bank_name,
		"promo_type": promo_type,
		"bonus_rate": bonus_rate,
		"category":   category,
	}
	_, err := m.C.InsertOne(ctx, insert)
	if err != nil {
		return err
	}
	return nil
}
