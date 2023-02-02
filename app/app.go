package app

import (
	"github.com/labstack/echo/v4"
)

type BookingInterface interface {
	BookingANumber(c echo.Context) error
	ListBooking(ctx echo.Context) error
}

type App struct {
	Booking BookingInterface
}

func (a *App) build() {
}

func NewApp(
	booking BookingInterface,
) *App {
	app := &App{
		Booking: booking,
	}
	app.build()

	return app
}
