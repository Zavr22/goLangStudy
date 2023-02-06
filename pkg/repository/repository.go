package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
