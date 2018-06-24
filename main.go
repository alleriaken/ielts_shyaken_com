package main

import (
	"ielts_study_backend/controllers/api/word"
	"github.com/joho/godotenv"
	"ielts_study_backend/models"
)


func main() {
	godotenv.Load()
	models.InitDB()
	routeApi()
	defer models.CloseDB()
}

func routeApi()  {
	word.HandleExam()
	return
}