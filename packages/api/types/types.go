package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	FirstName         string             `json:"first_name" bson:"firstName" validate:"required,min=2,max=30"`
	LastName          string             `json:"last_name" bson:"lastName" validate:"required,min=2,max=30"`
	UserName          string             `json:"user_name" bson:"userName" validate:"required,min=2,max=30"`
	Password          string             `json:"password" bson:"password"  validate:"required,min=6"`
	Email             string             `json:"email"    bson:"email"  validate:"email,required"`
	Phone             string             `json:"phone"  bson:"phone"    validate:"required"`
	Token             string             `json:"token" bson:"token"`
	Refresh_Token     string             `josn:"refresh_token" bson:"refreshToken"`
	CreatedAt         time.Time          `json:"created_at" bson:"createdAt"`
	UpdatedAt         time.Time          `json:"updtaed_at" bson:"updatedAt"`
	UserID            string             `json:"user_id" bson:"userId"`
	UserWalletAddress string             `json:"user_wallet_address" bson:"userWalletAddress"`
}

type Response struct {
	Message string `json:"message"`
	Status  int32  `json:"status"`
}
type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	OrderAmount string             `json:"orderAmount" bson:"orderAmount" validate:"required,min=2,max=30"`
	TotalPrice  string             `json:"totalPrice" bson:"totalPrice" validate:"required,min=2,max=30"`
}
