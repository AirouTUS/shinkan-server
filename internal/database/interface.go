package database

import (
	"github.com/AirouTUS/shinkan-server/internal/model"
)

var (
	ErrNotFound = "sql: no rows in result set"
)

type (
	ListCategoryInput struct{}
	GetCircleInput    struct {
		ID int
	}
	ListCirclesCircleTypesInput struct {
		ID int
	}
	ListCircleInput struct {
		CategoryID []int
	}
)

type DBRepository interface {
	ListCategory(input ListCategoryInput) ([]*model.Category, error)
	GetCircle(input GetCircleInput) (*model.Circle, error)
	ListCirclesCircleTypes(input ListCirclesCircleTypesInput) ([]*model.CirclesCircleTypes, error)
	ListCircle(input ListCircleInput) ([]*model.Circle, error)
}
