package service

import (
	goProj "github.com/Zavr22/goLangStudy"
	"github.com/Zavr22/goLangStudy/pkg/repository"
)

type ToDoListService struct {
	repo repository.ToDoList
}

func NewToDoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{repo: repo}
}

func (s *ToDoListService) Create(userId int, list goProj.ToDoList) (int, error) {
	return s.repo.Create(userId, list)
}
