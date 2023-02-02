package di

import (
	"context"

	"eh-digital-shift/app"
	"eh-digital-shift/app/handler"
	"eh-digital-shift/app/usecases"
	"eh-digital-shift/config"
	"eh-digital-shift/httpServer"
	"eh-digital-shift/services"
	"github.com/labstack/echo/v4"
)

func serverProvider(e *echo.Echo, cfg *config.Config, app *app.App) *httpServer.Server {
	return httpServer.NewServer(e, cfg, app)
}

func bookingANumberProvider(mongoService usecases.MongoServiceInterface) handler.BookingANumberUseCase {
	return usecases.NewBookingANumber(mongoService)
}

func bookingListProvider(mongoService usecases.MongoServiceInterface) handler.ListBookingUseCase {
	return usecases.NewListBookings(mongoService)
}

func bookingHandlerProvider(
	bookingANumber handler.BookingANumberUseCase, bookingList handler.ListBookingUseCase,
) app.BookingInterface {
	return handler.NewBooking(bookingANumber, bookingList)
}

func appProvider(booking app.BookingInterface) *app.App {
	return app.NewApp(booking)
}

func mongoServiceProvider(cfg *config.Config) usecases.MongoServiceInterface {
	mongo, err := services.NewMongoDB(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return mongo
}
