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
	GetAll(userId int) ([]goProj.ToDoList, error)
	GetById(userId int, listId int) (goProj.ToDoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input goProj.UpdateListInput) error
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
