package webserver

import (
	"github.com/asepnur/iskandar/src/webserver/handler"
	"github.com/julienschmidt/httprouter"
)

func loadRouter(r *httprouter.Router) {
	r.GET("/users", handler.SelectUserHandler)
	r.GET("/filter", handler.SelectUserFilterHandler)
}
