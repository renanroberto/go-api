package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", Logger(http.FileServer(http.Dir("./view"))))

	for _, route := range routes {
		router.
			Path(route.Path).
			Methods(route.Method).
			Handler(Logger(route.Handler))
	}

	return router
}
