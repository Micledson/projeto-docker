package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"projeto-docker/src/apps/api/handlers/dto/response"
	"projeto-docker/src/core/domain/errors"
	"strconv"
)

const COOKIE_TOKEN_NAME = "go_backend_template_v1"

func GetIDFromQueryParam(queryParam string, context echo.Context) (*uuid.UUID, error) {
	var uid *uuid.UUID
	idString := context.QueryParam(queryParam)
	if idString != "" {
		id, err := uuid.Parse(idString)
		if err != nil {
			valErr := errors.NewValidation([]string{fmt.Sprintf("'%s' inválido", queryParam)})
			return nil, response.ErrorBuilder().NewFromDomain(valErr)
		}

		uid = &id
	}
	return uid, nil
}

func GetIDFromPathParam(pathParam string, context echo.Context) (*uuid.UUID, error) {
	var uid *uuid.UUID
	idString := context.Param(pathParam)
	if idString != "" {
		id, err := uuid.Parse(idString)
		if err != nil {
			valErr := errors.NewValidation([]string{fmt.Sprintf("'%s' inválido", pathParam)})
			return nil, response.ErrorBuilder().NewFromDomain(valErr)
		}

		uid = &id
	}
	return uid, nil
}

func GetIntFromQueryParam(queryParam string, context echo.Context) (*int, error) {
	var code *int
	codeString := context.QueryParam(queryParam)
	if codeString != "" {
		parsedValue, err := strconv.Atoi(codeString)
		if err != nil {
			valErr := errors.NewValidation([]string{fmt.Sprintf("'%s' inválido", queryParam)})
			return nil, response.ErrorBuilder().NewFromDomain(valErr)
		}
		code = &parsedValue
	}
	return code, nil
}

func GetIsActiveFromQueryParam(queryParam string, context echo.Context) *bool {
	var isActive *bool
	isActiveString := context.QueryParam(queryParam)
	parsedValue, err := strconv.ParseBool(isActiveString)
	if err == nil {
		isActive = &parsedValue
	}

	return isActive
}
