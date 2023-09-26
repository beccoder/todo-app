package repository

import (
	"github.com/jmoiron/sqlx"
	todo_app "todo-app"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GetUser(username, password string) (todo_app.User, error)
}

type TodoList interface {
	Create(userId int, list todo_app.TodoList) (int, error)
	GetAll(userId int) ([]todo_app.TodoList, error)
	GetById(userId, listId int) (todo_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type Repository struct {
	Authorization
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
