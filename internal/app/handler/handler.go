package handler

import (
	"net/http"
	"strconv"

	"github.com/AirouTUS/shinkan-server/internal/usecase"

	"github.com/AirouTUS/shinkan-server/internal/app/input"

	"github.com/AirouTUS/shinkan-server/internal/model"

	"github.com/AirouTUS/shinkan-server/internal/app/output"
	"github.com/AirouTUS/shinkan-server/internal/database"
	"github.com/labstack/echo/v4"
)

var (
	Categories []*model.Category
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
	return APIResponseOK(c, output.ToCategoryList(Categories))
}

func (h *Handler) GetCircle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}
	if id <= 0 {
		return APIResponse(c, http.StatusBadRequest, "Bad Request")
	}

	circle, err := h.dbRepository.GetCircle(database.GetCircleInput{ID: id})
	if err != nil {
		if err.Error() == database.ErrNotFound {
			return APIResponse(c, http.StatusNotFound, "サークルが存在しません")
		}
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}

	circleTypes, err := h.dbRepository.ListCirclesCircleTypes(database.ListCirclesCircleTypesInput{ID: id})
	if err != nil {
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	return APIResponseOK(c, output.ToGetCircle(circle, circleTypes, Categories))
}

func (h *Handler) ListCircle(c echo.Context) error {
	var param input.ListCircleInput
	if err := c.Bind(&param); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}
	if err := param.Validate(); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}

	input := database.ListCircleInput{
		CategoryID: param.CategoryID,
	}

	circles, err := h.dbRepository.ListCircle(input)
	if err != nil {
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}

	categories := usecase.ParseCircleCategory(Categories, param.CategoryID)

	return APIResponseOK(c, output.ToListCircle(circles, categories))
}
