package database

import "github.com/AirouTUS/shinkan-server/internal/model"

type DBRepository interface {
	ListCategory(input ListCategoryInput) ([]*model.Category, error)
}
