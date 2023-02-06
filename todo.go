package goProj

type ToDoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
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
