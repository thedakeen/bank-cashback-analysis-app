package mongodb

import (
	"bank-cashback-analysis/backend/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		"title":         title,
		"source_url":    source_url,
		"bank_name":     bank_name,
		"promo_type":    promo_type,
		"bonus_rate":    bonus_rate,
		"category_name": category,
	}
	_, err := m.C.InsertOne(ctx, insert)
	if err != nil {
		return err
	}
	return nil
}

func (m *PromoModel) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.C.Drop(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *PromoModel) GetAllPromos(filters Filters) ([]*models.Promotion, Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(int64(filters.PageSize))
	findOptions.SetSkip(int64(filters.Page-1) * int64(filters.PageSize))

	cursor, err := m.C.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer cursor.Close(ctx)

	var promos []*models.Promotion
	if err = cursor.All(ctx, &promos); err != nil {
		return nil, Metadata{}, err
	}

	totalRecords, err := m.C.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(int(totalRecords), filters.Page, filters.PageSize)

	return promos, metadata, nil
}

/////////////////////////// HALYK /////////////////////////

func (m *PromoModel) SaveShopToDB(shop models.Shop) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.C.InsertOne(ctx, shop)
	return err
}

func (m *PromoModel) SavePromotionToDB(promotion models.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.C.InsertOne(ctx, promotion)
	return err
}
