package httpServer

import (
	"fmt"

	"eh-digital-shift/app"
	"eh-digital-shift/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Server *echo.Echo
	cfg    *config.Config
	app    *app.App
}

// Start run the server
func (s *Server) Start() {
	s.Server.Logger.Fatal(s.Server.Start(fmt.Sprintf(":%s", s.cfg.Port)))
}

func NewServer(
	server *echo.Echo,
	cfg *config.Config,
	app *app.App,
) *Server {
	return &Server{
		Server: server,
		cfg:    cfg,
		app:    app,
	}
}
