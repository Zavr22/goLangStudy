package main

import (
	"context"
	"github.com/Zavr22/goLangStudy/cmd/pkg/postgresql"
	User2 "github.com/Zavr22/goLangStudy/internal/User"
	User "github.com/Zavr22/goLangStudy/internal/User/db"
	"net/http"

	"log"
)

//goland:noinspection ALL
func main() {

	mux := http.NewServeMux()
	
	log.Println("Запуск веб-сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	client, err := postgresql.NewClient(context.TODO(), 10, "apple", "Rekbr345", "localhost", "5432", "postgres")
	if err != nil {
		log.Fatalf("Ошибка подключения к бд", err)
	}

	repository := User.NewRepository(client)

	newUser := User2.User{
		Name:     "Misha",
		Surname:  "Kulik",
		Password: "242",
		Role:     1,
	}
	err = repository.CreateUser(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newUser)
}
