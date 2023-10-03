package utils

import (
	"context"
	"errors"

	"todo_graphql/constants"

	"github.com/astaxie/beego/orm"
)

// HandleError is the method for returning the user facing error
func HandleError(c context.Context, errType string, err error) error {
	// req, _ := utils.RequestFromContext(c)
	// logger.Log.WithField("request", req).WithError(err).Error(errType)


	if errType != constants.InvalidRequestData {
		return errors.New(constants.ErrorString[errType])
	}
	if err == orm.ErrNoRows {
		err = errors.New(constants.NotFound)
	}

	errToReturn := constants.ErrorString[err.Error()]

	if errToReturn == "" {
		errToReturn = constants.ErrorString[errType]
	}

	return errors.New(errToReturn)

}
