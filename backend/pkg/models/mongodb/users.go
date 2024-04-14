package mongodb

import (
	"bank-cashback-analysis/backend/pkg/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserModel struct {
	C               *mongo.Collection
	CardsCollection *mongo.Collection
}

func NewUserModel(usersCollection, cardsCollection *mongo.Collection) *UserModel {
	return &UserModel{
		C:               usersCollection,
		CardsCollection: cardsCollection,
	}
}

func (m *UserModel) IsEmailExists(email string) (bool, error) {
	var result models.User
	err := m.C.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *UserModel) CheckEmail(email string) error {

	exists, err := m.IsEmailExists(email)
	if err != nil {
		return err
	}

	if exists {
		return models.ErrDuplicateEmail
	}
	return nil
}
func (m *UserModel) SignUpComplete(email, name, surname, phone, address, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	filter := bson.M{"email": email}

	update := bson.M{
		"$set": bson.M{
			"name":           name,
			"surname":        surname,
			"phone":          phone,
			"address":        address,
			"hashedPassword": hashedPassword,
			"role":           "user",
			"created":        time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err = m.C.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (string, string, error) {

	var result models.User
	err := m.C.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", "", models.ErrInvalidCredentials
		} else {
			return "", "", err
		}
	}

	err = bcrypt.CompareHashAndPassword(result.HashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", models.ErrInvalidCredentials
		}
		return "", "", err
	}

	return result.ID.Hex(), "", nil
}

///////////////////////////////////////////////////////

func (m *UserModel) SetCard(userId primitive.ObjectID, card_number, card_type, bank_name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cardInsert := bson.M{
		"card_number": card_number,
		"card_type":   card_type,
		"bank_name":   bank_name,
		"user_id":     userId,
	}

	_, err := m.CardsCollection.InsertOne(ctx, cardInsert)

	filter := bson.M{"_id": userId}
	update := bson.M{
		"$push": bson.M{"cards": cardInsert},
	}

	_, err = m.C.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) GetUserInfo(userId primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"_id": userId}

	projection := bson.M{
		"email":   1,
		"name":    1,
		"surname": 1,
		"address": 1,
		"phone":   1,
		"cards":   1,
	}

	err := m.C.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return &user, nil
}
