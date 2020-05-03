package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AirouTUS/shinkan-server/internal/app/api/output"
	"github.com/AirouTUS/shinkan-server/internal/database"
	"github.com/AirouTUS/shinkan-server/internal/database/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/AirouTUS/shinkan-server/internal/model"
)

var (
	id            = 1
	name          = "test_name"
	about         = "test_about"
	catchCopy     = "test_catchCopy"
	cost          = "test_cost"
	location      = "test_location"
	workTime      = "test_workTime"
	membersNumber = "test_membersNumber"
	description   = "test_description"
	categoryID    = 1
	email         = "test@test.com"
	twitter       = "test_twitter"
	url           = "test_url"
	eyeCatch      = "test_eyeCatch"
	typeID        = 1
	typeName      = "test_typeName"
	imageURL      = "test_imageURL"
	updatedAt     = "2020-10-10"
)

func TestMain(m *testing.M) {
	Categories = append(Categories, &model.Category{
		ID:   1,
		Name: "委員会",
	})
	code := m.Run()
	os.Exit(code)
}

func TestHandler_ListCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDBRepository(ctrl)

	e := echo.New()
	h := Handler{dbRepository: m}
	recContext := func() (*httptest.ResponseRecorder, echo.Context) {
		req := httptest.NewRequest(echo.GET, "/categories", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return rec, c
	}

	rec, c := recContext()
	if assert.NoError(t, h.ListCategory(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestHandler_GetCircle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDBRepository(ctrl)

	e := echo.New()
	h := Handler{dbRepository: m}

	recContext := func() (*httptest.ResponseRecorder, echo.Context) {
		req := httptest.NewRequest(echo.GET, "/circles/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		return rec, c
	}

	expectCircle := model.GetCircle{
		ID:            id,
		Name:          name,
		About:         about,
		CatchCopy:     catchCopy,
		Cost:          &cost,
		Location:      &location,
		WorkTime:      &workTime,
		MembersNumber: &membersNumber,
		Description:   description,
		CategoryID:    categoryID,
		Email:         email,
		Twitter:       twitter,
		URL:           url,
		EyeCatch:      eyeCatch,
		Types: []model.CircleType{
			{ID: typeID, Name: typeName},
		},
		Images: []model.CircleImages{
			{URL: imageURL},
		},
		UpdatedAt: updatedAt,
	}

	t.Run("success", func(t *testing.T) {
		m.EXPECT().GetCircle(
			database.GetCircleInput{
				ID: id,
			}).Return(&expectCircle, nil).AnyTimes()
		rec, c := recContext()
		if assert.NoError(t, h.GetCircle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var res output.Circle
			if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&res)) {
				assert.Equal(t, res.ID, id)
				assert.Equal(t, res.Name, name)
				assert.Equal(t, res.About, about)
				assert.Equal(t, res.CatchCopy, catchCopy)
				assert.Equal(t, res.Cost, cost)
				assert.Equal(t, res.Description, description)
				assert.Equal(t, res.EyeCatch, eyeCatch)
				assert.Equal(t, res.Twitter, twitter)
				assert.Equal(t, res.Email, email)
				assert.Equal(t, res.URL, url)
				for _, v := range res.Images {
					assert.Equal(t, v.URL, imageURL)
				}
				for _, v := range res.Types {
					assert.Equal(t, v.ID, typeID)
					assert.Equal(t, v.Name, typeName)
				}
				assert.Equal(t, res.Category.ID, categoryID)
				assert.Equal(t, res.UpdatedAt, updatedAt)
			}
		}
	})
}
