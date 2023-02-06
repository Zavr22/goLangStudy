package service

import "github.com/Zavr22/goLangStudy/pkg/repository"

type Authorization interface {
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
