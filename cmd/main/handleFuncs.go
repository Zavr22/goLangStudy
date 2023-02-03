package main

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var input user

}
