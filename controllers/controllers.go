package controllers


import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love English")
}