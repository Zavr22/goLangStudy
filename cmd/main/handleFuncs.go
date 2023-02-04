package main

import (
	"context"
	"fmt"
	"github.com/Zavr22/goLangStudy/internal/User"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func login(c *gin.Context, q User.Storage) {
	type user struct {
		Name     string `json:"name" `
		Password string `json:"password" `
	}
	var input user
	if err := c.BindJSON(&input); err != nil {
		fmt.Print(c, http.StatusBadRequest, err.Error())
		return
	}
	var i bool
	data1, _ := c.GetPostForm("name")
	var data2 User.User
	data2, err := q.FindOne(context.Background(), data1)
	if err != nil {
		log.Fatal("error")
	}

	if data1 == data2.Name {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok", // cast it to string before showing
		})
	}

	if i == true {

	}

}
