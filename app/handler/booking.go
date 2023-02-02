package handler

import (
	"context"
	"net/http"

	"eh-digital-shift/app/dto"
	"github.com/labstack/echo/v4"
)

type BookingANumberUseCase interface {
	Invoke(ctx context.Context, booking *dto.Booking) error
}
type ListBookingUseCase interface {
	Invoke(ctx context.Context) ([]dto.Booking, error)
}

type Booking struct {
	BookingANumberUseCase BookingANumberUseCase
	ListBookingUseCase    ListBookingUseCase
}

func (b *Booking) BookingANumber(ctx echo.Context) error {
	bookingDto := &dto.Booking{}

	if err := ctx.Bind(bookingDto); err != nil {
		ctx.Error(err)
		return nil
	}

	err := b.BookingANumberUseCase.Invoke(ctx.Request().Context(), bookingDto)
	if err != nil {
		r := dto.Response{Message: err.Error()}

		return ctx.JSON(http.StatusBadRequest, r)
	}

	return ctx.JSON(http.StatusOK, dto.Response{Message: "OK"})
}

func (b *Booking) ListBooking(ctx echo.Context) error {
	bookings, err := b.ListBookingUseCase.Invoke(ctx.Request().Context())

	if err != nil {
		r := dto.Response{Message: err.Error()}

		return ctx.JSON(http.StatusBadRequest, r)
	}

	return ctx.JSON(http.StatusOK, bookings)
}

func NewBooking(bookingANumber BookingANumberUseCase, bookingList ListBookingUseCase) *Booking {
	return &Booking{BookingANumberUseCase: bookingANumber, ListBookingUseCase: bookingList}
}
