package goProj

import "errors"

type ToDoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	id     int
	UserId int
	ListId int
}

type ToDoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ToDo        bool   `json:"done"`
}

type ListsItem struct {
	id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
