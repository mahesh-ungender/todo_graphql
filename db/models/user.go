package models

import (
	graphmodel "todo_graphql/graph/model"

	"github.com/astaxie/beego/orm"
)

// todo model for todo
type Todo struct {
	ID       int64   `orm:"pk;auto;column(id)" json:"id"`
	ItemName *string `json:"itemname" orm:"null"`
	Status   *string `json:"status" orm:"null"`
}

// TableName specifies the table name for User in the db
func (t *Todo) TableName() string {
	return "items"
}

// Serialize serializes the user model to response
func (t *Todo) Serialize() *graphmodel.Todo {
	res := &graphmodel.Todo{
		ID:       string(t.ID),
		ItemName: *t.ItemName,
		Status:   *t.Status,
	}
	return res
}

// register the tablw
func init() {
	orm.RegisterModel(new(Todo))
}
