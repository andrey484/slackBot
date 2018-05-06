package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"webhook",
		"POST",
		"/webhook",
		slackWebHook,
	},
}

func NewRouter() *mux.Router  {
	router := mux.NewRouter().StrictSlash(false)

	for _, route := range routes{

		var handler http.Handler
		handler = route.HandleFunc
		handler = MyLogger(handler, route.Name)

		router.Methods(route.Method).Name(route.Name).Path(route.Pattern).HandlerFunc(route.HandleFunc)
	}

	return router
}

