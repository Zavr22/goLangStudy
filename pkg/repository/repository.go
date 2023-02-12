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
	}
}
