package handlers

import (
	"endpointsService/repositories"
	"endpointsService/structures"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateEndpoint(ctx echo.Context, endpointsRepository *repositories.EndpointRepository) error {
	body := new(structures.UpdateBodyDTO)

	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "id is required"})
	}

	endpoint, getErr := endpointsRepository.GetById(ctx.Request().Context(), id)

	if getErr != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": getErr.Error()})
	}

	endpoint.UpdateFromBody(body)

	err := endpointsRepository.Update(ctx.Request().Context(), endpoint)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, endpoint)
}
