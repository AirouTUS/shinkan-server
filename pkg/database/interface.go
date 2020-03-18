package database

import (
	"github.com/AirouTUS/shinkan-server/pkg/model"
)

var (
	ErrNotFound = "sql: no rows in result set"
)

type (
	// api
	ListCategoryInput struct{}
	GetCircleInput    struct {
		ID int
	}
	ListCircleInput struct {
		CategoryID []int
	}

	// admin
	PostCircleInput struct {
		Name        string
		About       string
		CatchCopy   string
		Description string
		EyeCatch    string
		Email       string
		Twitter     string
		URL         string
		Images      []InputCircleImage
		Types       []InputCircleType
		CategoryID  int
	}
)

type DBRepository interface {
	// api
	ListCategory(input ListCategoryInput) ([]*model.Category, error)
	GetCircle(input GetCircleInput) (*model.GetCircle, error)
	ListCircle(input ListCircleInput) ([]*model.Circle, error)

	// admin
	PostCircle(input PostCircleInput) error
}
