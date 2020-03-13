package handler

import (
	"net/http"

	"github.com/AirouTUS/shinkan-server/internal/app/output"
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
	categories, err := h.dbRepository.ListCategory(database.ListCategoryInput{})
	if err != nil {
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	return APIResponseOK(c, output.ToCategoryList(categories))
}
