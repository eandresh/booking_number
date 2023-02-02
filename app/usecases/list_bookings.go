package usecases

import (
	"context"

	"eh-digital-shift/app/dto"
	"go.mongodb.org/mongo-driver/bson"
)

type ListBookings struct {
	mongo MongoServiceInterface
}

func (l *ListBookings) Invoke(ctx context.Context) ([]dto.Booking, error) {
	bookings, err := l.mongo.FindAll(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	booking := dto.Booking{}
	result, _ := booking.ToMapFromMongoResult(bookings)

	return result, nil
}

func NewListBookings(mongo MongoServiceInterface) *ListBookings {
	return &ListBookings{
		mongo: mongo,
	}
}
