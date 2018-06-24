package controllers

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Message string `json:"msg"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
	//= {
	// 		"msg":"Hi",
	//		"code":200,
	//		"data": null
	// }
}

func Index(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&Response{Message: "This is shyaken's ielts study framework"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}