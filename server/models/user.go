package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBUser struct {
	ID             primitive.ObjectID `bson:"_id"`
	Email          *string            `json:"email,omitempty"`
	HashedPassword *string            `json:"hashed_password,omitempty"`
}
