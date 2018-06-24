package main

import (
	"github.com/joho/godotenv"
	"ielts_study_backend/models"
	"net/http"
	"fmt"
	"os"

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
	defer models.CloseDB()
}
