package handler

import (
	"net/http"

	"github.com/AirouTUS/shinkan-server/pkg/app/admin/input"

	"github.com/AirouTUS/shinkan-server/pkg/model"
	"github.com/labstack/echo/v4"

	"github.com/AirouTUS/shinkan-server/pkg/database"
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

func (h *Handler) PostCircle(c echo.Context) error {
	var param input.PostCircleInput
	if err := c.Bind(&param); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}
	if err := param.Validate(); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}

	input := database.PostCircleInput{
		Name:        param.Name,
		About:       param.About,
		CatchCopy:   param.CatchCopy,
		Description: param.Description,
		EyeCatch:    param.EyeCatch,
		Email:       param.Email,
		Twitter:     param.Twitter,
		URL:         param.URL,
		CategoryID:  param.CategoryID,
	}
	input.Types = make([]database.InputCircleType, 0, len(param.Types))
	for _, v := range param.Types {
		input.Types = append(input.Types, database.InputCircleType(v))
	}
	input.Images = make([]database.InputCircleImage, 0, len(param.Images))
	for _, v := range param.Images {
		input.Images = append(input.Images, database.InputCircleImage(v))
	}

	err := h.dbRepository.PostCircle(input)
	if err != nil {
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	return APIResponse(c, http.StatusCreated, "")
}
