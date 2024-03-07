package http

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRoute(app *server.Hertz, handler *Handler) {
	app.GET("/ping", handler.Ping)

}
