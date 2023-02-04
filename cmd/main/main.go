package main

import (
	"context"
	"github.com/Zavr22/goLangStudy/cmd/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
)

//goland:noinspection ALL
func main() {

	mux := http.NewServeMux()

	log.Println("Запуск веб-сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	_, err = postgresql.NewClient(context.TODO(), 10, "apple", "Rekbr345", "localhost", "5432", "postgres")
	if err != nil {
		log.Fatalf("Ошибка подключения к бд", err)
	}
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/addUser", )
	}

	/*newUser := User2.User{
		Name:     "Misha",
		Surname:  "Kulik",
		Password: "242",
		Role:     1,
	}
	err = repository.CreateUser(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newUser)*/

}
