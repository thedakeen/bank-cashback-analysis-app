package mongodb

import (
	"bank-cashback-analysis/backend/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PromoModel struct {
	C *mongo.Collection
}

func NewPromotionModel(c *mongo.Collection) *PromoModel {
	return &PromoModel{
		C: c,
	}
}

func (m *PromoModel) SavePromotionToDB(promotion models.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.C.InsertOne(ctx, promotion)
	return err
}
