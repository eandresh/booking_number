package handler

import (
	"context"

	"eh-digital-shift/app/dto"
	"github.com/stretchr/testify/mock"
)

type bookingANumberUseCaseMock struct {
	mock.Mock
}

func (banum *bookingANumberUseCaseMock) Invoke(ctx context.Context, booking *dto.Booking) error {
	args := banum.Called(ctx, booking)

	return args.Error(0)
}

type listBookingUseCaseMock struct {
	mock.Mock
}

func (lbum *listBookingUseCaseMock) Invoke(ctx context.Context) ([]dto.Booking, error) {
	args := lbum.Called(ctx)

	return args.Get(0).([]dto.Booking), args.Error(1)
}
