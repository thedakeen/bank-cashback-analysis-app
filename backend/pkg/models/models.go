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