package main

import (
	"github.com/joho/godotenv"
	"ielts_study_backend/models"
	"net/http"
	"fmt"
	"os"
	"ielts_study_backend/controllers"
	"ielts_study_backend/controllers/api/word"
	"github.com/gorilla/mux"
)

func main() {
	godotenv.Load()
	models.InitDB()
	routeApi()
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", os.Getenv("SERVE_HOST"), os.Getenv("SERVE_PORT")), nil);
		err != nil {
		panic(err)
	}
}

type RouteElement struct {
	Methods []string
	Handler interface{}
}

func Route(handler interface{}, method []string) RouteElement {
	if len(method) == 0  {
		method = []string{"GET", "POST", "PUT"} // default handler will serve 3 methods: GET, POST, PUT
	}
	fmt.Println(method)
	return RouteElement{method, handler}  // enforce the default value here
}

func routeApi()  {

	routes := map[string]RouteElement{
		"/" : Route(controllers.Index, []string{}),

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