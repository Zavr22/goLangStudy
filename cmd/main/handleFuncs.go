package main

import (
	"context"
	"fmt"
	"github.com/Zavr22/goLangStudy/cmd/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func login(c *gin.Context, ctx *context.Context) {
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
	var data2 string
	rows, err := postgresql.Client.Query(ctx, `select USER.NAME FROM PUBLI.USER WHERE USER.NAME = $1`, data1)
	if err != nil {
		log.Fatal("error")
	}
	data2 = rows
	if data1 == data2 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok", // cast it to string before showing
		})
	}

	if i == true {

	}

}
