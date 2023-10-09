package repository

import (
	"context"
	"todo_graphql/db/models"
	"todo_graphql/logger"

	"github.com/astaxie/beego/orm"
)

type TodoRepo interface {
	Save(ctx context.Context, doc *models.Todo ) error
	GetAll(ctx context.Context, query models.Todo) ([]*models.Todo,int64, error)
	FindByID(ctx context.Context, id int64) (*models.Todo, error)
	Update(ctx context.Context, doc *models.Todo, fieldsToUpdate []string) error
	Delete(ctx context.Context, ID int64) error
}

type todoRepo struct {
	db        orm.Ormer
}

// Save saves the item in the database
func (repo *todoRepo) Save(ctx context.Context, doc *models.Todo) error {
	groupError := "CREATE_ITEM"

	logger.Log.Info("Inserting item in the db")
	
	id, err := repo.db.Insert(doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("Done inserting item in the db")
	doc.ID = id
	return nil
}

// GetAll returns all the users in the db
func (repo *todoRepo) GetAll(ctx context.Context, query models.Todo) ([]*models.Todo,int64, error) {
	groupError := "GET_ALL_ITEMS"
	var items []*models.Todo

	logger.Log.Info("Getting all items")
	qs := repo.db.QueryTable(new(models.Todo))

	// num := qs.Distinct().All(&items)
	// logger.Log.Info(fmt.Sprintf("Read %d items from the db", num))

	cnt, err := qs.GroupBy("id").Count()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return items, 0, err
	}

	return items, cnt, nil
}

// Update updates the existing item's attributes in the database
func (repo *todoRepo) Update(ctx context.Context, doc *models.Todo, fieldsToUpdate []string) error {
	groupError := "UPDATE_ITEM"

	//doc.UpdatedAt = time.Now().Unix()

			logger.Log.Info("Updating item in the db")
	_, err := repo.db.Update(doc, fieldsToUpdate...)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("Done updating item in the db")

	return nil
}

// Delete a item in the db
func (repo *todoRepo) Delete(ctx context.Context, ID int64) error {
	groupError := "DELETE_ITEM"
	user := models.Todo{
		ID:         ID,
	}

	logger.Log.Info("Deleting item in the db")
	_, err := repo.db.Delete(&user)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("Done deleting item in the db")
	return nil
}


// FindByID returns the user associated with the id parameter
func (repo *todoRepo) FindByID(ctx context.Context, id int64) (*models.Todo, error) {
	groupError := "FIND_ITEM_BY_ID"
	todo := models.Todo{}

	logger.Log.Info("Finding item by id in the db")

	qs := repo.db.QueryTable(new(models.Todo))
	err := qs.Filter("id", id).Filter("IsInactive__exact", false).One(&todo)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return &todo, err
	}
	logger.Log.Info("Done finding item by id in the db")

	return &todo, nil
}



func NewTodoRepo(db orm.Ormer) TodoRepo {
	return &todoRepo{
		db:        db,
	}
}
