package service

import (
	"todo_graphql/api"
	"todo_graphql/instance"
	"todo_graphql/repository"
)

// Services is the interface for enclosing all the services
type Services interface {
	Todo() api.Todo
}

type services struct {
	todoService            api.Todo
}


func (svc *services) Todo() api.Todo {
	return svc.todoService
}

// Init initializes the services
func Init() Services {
	db := instance.DB()

	return &services{
		todoService: api.NewTodo(
			repository.NewTodoRepo(db),
		),
	}
}
