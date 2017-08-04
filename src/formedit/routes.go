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

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"SimpleForm",
		"GET",
		"/",
		SimpleForm,
	},
        Route{
              	"SimpleEdit",
                "GET",
                "/edit",
                SimpleEdit,
        },
		Route{
              	"SaveEdit",
                "post",
                "/save",
                SaveEdit,
        },
	        Route{
                "Download",
               	"get",
                "/download/{id}",
                Download,
        },
          	Route{
                "InsertView",
                "get",
                "/insertview/{id}",
                InsertView,
        },
          	Route{
                "TableInsertView",
                "get",
                "/tableinsertview/{id}",
                TableInsertView,
        },

}

