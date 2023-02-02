package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID       primitive.ObjectID `bson:"_id"`
	ClientID string             `bson:"client_id"`
	Number   int                `bson:"number"`
	CreateAt time.Time          `bson:"create_at"`
	UpdateAt time.Time          `bson:"update_at"`
}
