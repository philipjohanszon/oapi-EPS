package handlers

import (
	"endpointsService/repositories"
	"endpointsService/structures"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateEndpoint(ctx echo.Context, endpointsRepository *repositories.EndpointRepository) error {
	body := new(structures.CreateBody)

	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	endpoint, err := endpointsRepository.Create(ctx.Request().Context(), body)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, endpoint)
}
