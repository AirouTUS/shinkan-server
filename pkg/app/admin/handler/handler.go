package handler

import (
	"net/http"
	"strconv"

	"github.com/AirouTUS/shinkan-server/pkg/usecase"

	"github.com/AirouTUS/shinkan-server/pkg/app/admin/input"

	"github.com/AirouTUS/shinkan-server/pkg/model"

	"github.com/AirouTUS/shinkan-server/pkg/app/api/output"
	"github.com/AirouTUS/shinkan-server/pkg/database"
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

	circles, err := h.dbRepository.GetCircle(database.GetCircleInput{ID: id})
	if err != nil {
		if err.Error() == database.ErrNotFound {
			return APIResponse(c, http.StatusNotFound, "サークルが存在しません")
		}
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	if len(circles) <= 0 {
		return APIResponse(c, http.StatusNotFound, "サークルが存在しません")
	}

	return APIResponseOK(c, output.ToGetCircle(circles, Categories))
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
		if err.Error() == database.ErrNotFound {
			return APIResponse(c, http.StatusNotFound, "サークルが存在しません")
		}
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	if len(circles) <= 0 {
		return APIResponse(c, http.StatusNotFound, "サークルが存在しません")
	}
	result := usecase.ParseCircles(circles, Categories)

	return APIResponseOK(c, output.ToListCircle(result))
}
