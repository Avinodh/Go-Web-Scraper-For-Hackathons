package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},

	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},

	Route{
		"FetchBlogs",
		"GET",
		"/api/blogs",
		FetchBlogs,
	},

	Route{
		"PostBlog",
		"POST",
		"/api/blogs",
		PostBlog,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		/*********** LOGGER CODE *************/
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		/*************************************/

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler) //Analogous to Handler(route.handlerFunc)
	}

	s := http.StripPrefix("/", http.FileServer(http.Dir("./")))
	router.PathPrefix("/").Handler(s)

	return router
}