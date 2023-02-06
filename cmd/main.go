package main

import (
	goProj "github.com/Zavr22/goLangStudy"
	"github.com/Zavr22/goLangStudy/pkg/handler"
	"github.com/Zavr22/goLangStudy/pkg/repository"
	"github.com/Zavr22/goLangStudy/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(goProj.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server : %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
