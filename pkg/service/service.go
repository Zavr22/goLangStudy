package service

import (
	goProj "github.com/Zavr22/goLangStudy"
	"github.com/Zavr22/goLangStudy/pkg/repository"
)

type Authorization interface {
	CreateUser(user goProj.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list goProj.ToDoList) (int, error)
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      NewToDoListService(repos.ToDoList),
	}
}
