package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_app.TodoList) (int, error)
	GetAll(userId int) ([]todo_app.TodoList, error)
	GetById(userId, listId int) (todo_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type Service struct {
	Authorization
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
