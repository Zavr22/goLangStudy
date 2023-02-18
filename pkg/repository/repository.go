package repository

import (
	goProj "github.com/Zavr22/goLangStudy"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user goProj.User) (int, error)
	GetUser(username, password string) (goProj.User, error)
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

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewTodoListPostgres(db),
	}
}
