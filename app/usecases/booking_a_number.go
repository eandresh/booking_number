package usecases

import (
	"context"
	"errors"
	"fmt"

	"eh-digital-shift/app/dto"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoServiceInterface interface {
	Insert(ctx context.Context, obj interface{}) error
	FindAll(ctx context.Context, filter interface{}) ([]bson.M, error)
}

type BookingANumber struct {
	mongo MongoServiceInterface
}

func (ban *BookingANumber) Invoke(ctx context.Context, booking *dto.Booking) error {
	if !ban.validateNumber(ctx, booking) {
		return errors.New("ERROR - Número ya reservado")
	}

	if !ban.validateClient(ctx, booking) {
		return errors.New(fmt.Sprintf("ERROR - Cliente %s ya tiene reserva", booking.ClientID))
	}

	err := ban.mongo.Insert(ctx, booking)
	if err != nil {
		return err
	}

	return nil
}

func (ban *BookingANumber) validateNumber(ctx context.Context, booking *dto.Booking) bool {
	filter := bson.D{{"number", booking.Number}}

	bookings, err := ban.mongo.FindAll(ctx, filter)

	if err != nil {
		return false
	}

	return len(bookings) == 0
}

func (ban *BookingANumber) validateClient(ctx context.Context, booking *dto.Booking) bool {
	filter := bson.D{{"client_id", booking.ClientID}}
	bookings, err := ban.mongo.FindAll(ctx, filter)

	if err != nil {
		return false
	}

	return len(bookings) == 0
}

func NewBookingANumber(mongo MongoServiceInterface) *BookingANumber {
	return &BookingANumber{
		mongo: mongo,
	}
}
