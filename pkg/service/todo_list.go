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

func (s *ToDoListService) GetAll(userId int) ([]goProj.ToDoList, error) {
	return s.repo.GetAll(userId)
}

func (s *ToDoListService) GetById(userId int, listId int) (goProj.ToDoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ToDoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *ToDoListService) Update(userId, listId int, input goProj.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
