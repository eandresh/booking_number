//go:build wireinject
// +build wireinject

package di

import (
	"eh-digital-shift/httpServer"
	"github.com/google/wire"
)

func Initialize() (*httpServer.Server, error) {
	wire.Build(stdSet)

	return &httpServer.Server{}, nil
}
