package handlers

import (
	"endpointsService/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	equals      = "=="
	notEquals   = "!="
	lessThan    = "<"
	greaterThan = ">"
)

const (
	projectId = "projectId"
	url       = "url"
	method    = "method"
)

func GetEndpoints(ctx echo.Context, endpointsRepository *repositories.EndpointRepository) error {
	parameters := getParameters(ctx)

	endpoints, err := endpointsRepository.Get(ctx.Request().Context(), parameters)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, endpoints)
}

func getParameters(ctx echo.Context) []repositories.QueryParameter {
	queryParameters := make([]repositories.QueryParameter, 0)

	projectIdValue := ctx.QueryParam("projectId")

	urlValue := ctx.QueryParam("url")
	methodValue := ctx.QueryParam("method")

	if projectIdValue != "" {
		param := repositories.NewQueryParameter(projectId, equals, projectIdValue)

		queryParameters = append(queryParameters, param)
	}

	if urlValue != "" {
		param := repositories.NewQueryParameter(url, equals, urlValue)

		queryParameters = append(queryParameters, param)
	}

	if methodValue != "" {
		param := repositories.NewQueryParameter(method, equals, methodValue)

		queryParameters = append(queryParameters, param)
	}

	return queryParameters
}
