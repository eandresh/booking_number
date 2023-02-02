package main

import (
	"eh-digital-shift/di"
)

func main() {
	server, err := di.Initialize()
	if err != nil {
		panic("main:: fatal err getting handle: " + err.Error())
	}
	server.Routes()
	server.Start()
}
