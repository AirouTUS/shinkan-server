package database

import "github.com/AirouTUS/shinkan-server/internal/model"

type CategoryList []Category

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (m CategoryList) category() []*model.Category {
	result := make([]*model.Category, 0, len(m))
	for _, v := range m {
		content := model.Category{
			ID:   v.ID,
			Name: v.Name,
		}
		result = append(result, &content)
	}
	return result
}
