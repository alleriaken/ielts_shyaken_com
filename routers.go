package main

import (
	"ielts_study_backend/controllers/api/word"
	"github.com/gorilla/mux"
	"net/http"
	"ielts_study_backend/controllers"
	"ielts_study_backend/controllers/api"
)

type RouteElement struct {
	Methods []string
	Handler interface{}
}

func Route(handler interface{}, method []string) RouteElement {
	if len(method) == 0  {
		method = []string{"GET", "POST", "PUT"} // default handler will serve 3 methods: GET, POST, PUT
	}
	return RouteElement{method, handler}  // enforce the default value here
}

func routeApi()  {

	routes := map[string]RouteElement{
		"/" : Route(controllers.Index, []string{}),
		"/api" : Route(api.Index, []string{}),

		// Routes for word apis
		"/api/word" : Route(word.HandleGet, []string{}),
	}

	router := mux.NewRouter()

	for route, handler := range routes  {
		router.HandleFunc(route, handler.Handler.(func(w http.ResponseWriter, r *http.Request))).Methods(handler.Methods...)
	}

	http.Handle("/", router)
	return
}