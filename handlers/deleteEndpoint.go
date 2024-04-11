package handlers

import (
	"endpointsService/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteEndpoint(ctx echo.Context, endpointsRepository *repositories.EndpointRepository) error {
	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "id required"})
	}

	err := endpointsRepository.Delete(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
