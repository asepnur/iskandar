package webserver

import (
	"github.com/asepnur/iskandar/src/webserver/handler"
	"github.com/julienschmidt/httprouter"
)

func loadRouter(r *httprouter.Router) {
	r.GET("/", handler.TestingHTML)
	r.GET("/name", handler.ViewHTML)
	r.GET("/users", handler.SelectUserHandler)
}
