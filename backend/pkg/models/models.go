package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrInvalidCredentials = errors.New("models: invalid credentials")
var ErrDuplicateEmail = errors.New("models: duplicate email")
var ErrEmailDoesNotExist = errors.New("models: email does not exist")

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Email          string             `bson:"email"`
	Name           string             `bson:"name"`
	Surname        string             `bson:"surname"`
	HashedPassword []byte             `bson:"hashedPassword"`
	Created        time.Time          `bson:"created"`
	Phone          string             `bson:"phone"`
	Address        string             `bson:"address"`
	Role           string             `bson:"role"`
	OTP            OTP                `bson:"otp,omitempty"`
}

type OTPs struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	OTP      OTP                `bson:"otp,omitempty"`
	Verified bool               `bson:"verified"`
}
type OTP struct {
	Code    string    `bson:"code"`
	Expires time.Time `bson:"expires"`
}

type Cards struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Number   string             `bson:"number"`
	CardType string             `bson:"card_type"`
	BankName string             `bson:"bank_name"`
}

type Promotion struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"title"`
	SourceUrl    string             `bson:"source_url"`
	BankName     string             `bson:"bank_name"`
	CompanyName  string             `bson:"company_name"`
	CategoryName string             `bson:"category_name"`
	BonusRate    float64            `bson:"bonus_rate"`
	Requirements string             `bson:"requirements"`
	Restrictions string             `bson:"restrictions"`
	Expires      time.Time          `bson:"expires-at"`
	Start        time.Time          `bson:"start-at"`
	Location     string             `bson:"location"`
	Type         string             `bson:"promo_type"`
	CardType     string             `bson:"card_type"`
}
