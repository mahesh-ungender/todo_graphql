package api

import (
	"context"
	"errors"
	"todo_graphql/constants"
	"todo_graphql/db/models"
	graphmodel "todo_graphql/graph/model"
	"todo_graphql/utils"

	apiutils "todo_graphql/api/utils"
	"todo_graphql/repository"

	"github.com/astaxie/beego/orm"
)

type Todo interface {
	Create(ctx context.Context, data graphmodel.Todo) (*graphmodel.NewTodo, error)
	GetAllItems(ctx context.Context, input graphmodel.Todo) (*graphmodel.TodoList, error)
	UpdateItem(ctx context.Context, data graphmodel.Todo) (*graphmodel.Todo, error)
	RemoveItem(ctx context.Context, data graphmodel.Todo) (bool, error)
}

type todo struct {
	todoRepo            repository.TodoRepo
}

func (c *todo) Create(ctx context.Context, data graphmodel.Todo) (*graphmodel.NewTodo, error) {
	u := &graphmodel.NewTodo{}

	doc := &models.Todo{
		ItemName:       &data.ItemName,
		Status:        &data.Status,
	}

	err := c.todoRepo.Save(ctx, doc)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"item_duplicate\"" {
			return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.ITEM_ALREADY))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	
	// var todos []*models.Todo
	
	return u, nil
}

// GetAllUsers is the resolver for listing all the users
func (c *todo) GetAllItems(ctx context.Context, input graphmodel.Todo) (*graphmodel.TodoList, error) {
	todosList := graphmodel.TodoList{}

	//totalRows := c.todoRepo.GetAll(ctx)

	var u []*graphmodel.Todo

	todosList = graphmodel.TodoList{
		Todos:        u,
	}

	return &todosList, nil
}


// UpdateUser is the resolver for updating a user
func (c *todo) UpdateItem(ctx context.Context, data graphmodel.Todo) (*graphmodel.Todo, error) {
	var u *graphmodel.Todo

	doc, err := c.todoRepo.FindByID(ctx, data.ID)

	if err != nil {
		if err == orm.ErrNoRows {
			return u, apiutils.HandleError(ctx, constants.NotFound, err)
		}
		return u, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	// update entries
	doc.ItemName = utils.CheckNullAndSet(doc.ItemName, &data.ItemName)
	doc.Status = utils.CheckNullAndSet(doc.Status, &data.Status)

	err = c.todoRepo.Update(ctx,doc, []string{})
	if err != nil {
		if err == orm.ErrNoRows {
			return u, apiutils.HandleError(ctx, constants.NotFound, err)
		}
		return u, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	return u, nil
}


// RemoveUserFromTeam is the resolver for removing a user from a team
func (c *todo) RemoveItem(ctx context.Context, data graphmodel.Todo) (bool, error) {
	
	err := c.todoRepo.Delete(ctx, data.ID)
	
	if err != nil {
		if err == orm.ErrNoRows {
			return false, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.TeamDoesNotExist))
		}
		return false, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	return true, nil
}

// NewUser is the initialization method for the user resolvers
func NewTodo(todoRepo repository.TodoRepo,) Todo {
	return &todo{
		todoRepo:         todoRepo,
	}
}
