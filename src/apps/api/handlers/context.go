package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"projeto-docker/src/core/domain/errors"
	"strconv"
)

type RichContext interface {
	echo.Context

	AccountID() *uuid.UUID
	ProfileID() *uuid.UUID
	RoleName() string
	IsAdmin() bool

	GetUUIDPathParam(key string) (*uuid.UUID, errors.Error)
}

type RichHandler = func(RichContext) error

type richContext struct {
	echo.Context

	accountID *uuid.UUID
	profileID *uuid.UUID
	roleName  *string
}

func (c *richContext) RoleName() string {
	//TODO implement me
	panic("implement me")
}

func (c *richContext) IsAdmin() bool {
	//TODO implement me
	panic("implement me")
}

func NewRichContext(ctx echo.Context) (*richContext, error) {
	return &richContext{ctx, nil, nil, nil}, nil
}

func (c *richContext) AccountID() *uuid.UUID {
	return c.accountID
}

func (c *richContext) ProfileID() *uuid.UUID {
	return c.profileID
}

func (c *richContext) GetUUIDPathParam(key string) (*uuid.UUID, errors.Error) {
	strUUID := c.Param(key)
	if strUUID == "" {
		return nil, errors.NewFromString(fmt.Sprintf("you must provide a valid %s", key))
	} else if uuid, err := uuid.Parse(strUUID); err != nil {
		return nil, errors.NewFromString("the provided id is invalid")
	} else {
		return &uuid, nil
	}
}

func (c *richContext) GetStringPathParam(key string) string {
	return c.GetPathParam(key)
}

func (c *richContext) GetIntPathParam(key string) (int, *echo.HTTPError) {
	value := c.GetPathParam(key)
	if intValue, err := strconv.Atoi(value); err != nil {
		return 0, &echo.HTTPError{
			Message: fmt.Sprintf("the provided value for %s must be an integer", key),
		}
	} else {
		return intValue, nil
	}
}

func (c *richContext) GetPathParam(key string) string {
	value := c.Param(key)
	return value
}
