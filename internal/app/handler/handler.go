package handler

import (
	"net/http"

	"github.com/AirouTUS/shinkan-server/internal/app/input"
	"github.com/AirouTUS/shinkan-server/internal/database"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	dbRepository database.DBRepository
}

func NewHandler(
	dbRepository database.DBRepository,
) *Handler {
	return &Handler{
		dbRepository: dbRepository,
	}
}

func (h *Handler) ListCategory(c echo.Context) error {
	var param input.ListCategoryInput
	if err := c.Bind(&param); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}
	//categories, err := h.dbRepository.ListCategory(database.ListCategoryInput{})
	return APIResponseOK(c, param)
}
