package router

import (
	"net/http"

	"api/controller"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"/users", "GET", controller.GetUsers},
	Route{"/users", "POST", controller.CreateUser},
	Route{"/users/{id:[0-9]+}", "GET", controller.GetUserById},
	Route{"/users/{id:[0-9]+}", "PUT", controller.UpdateUser},
	Route{"/users/{id:[0-9]+}", "DELETE", controller.DeleteUser},
}
