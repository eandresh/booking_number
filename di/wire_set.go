package di

import (
	"eh-digital-shift/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var stdSet = wire.NewSet(
	echo.New,
	config.NewConfig,
	mongoServiceProvider,
	bookingANumberProvider,
	bookingListProvider,
	bookingHandlerProvider,
	appProvider,
	serverProvider,
)
