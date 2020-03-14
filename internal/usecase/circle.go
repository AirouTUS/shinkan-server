package usecase

import (
	"github.com/AirouTUS/shinkan-server/internal/model"
)

type CircleCategory struct {
	ID   int
	Name string
}

func ParseCircleCategory(categories []*model.Category, categoryIDs []int) (result []CircleCategory) {
	result = make([]CircleCategory, 0, len(categoryIDs))
	for _, v := range categoryIDs {
		for _, vv := range categories {
			if v == vv.ID {
				content := CircleCategory{
					ID:   vv.ID,
					Name: vv.Name,
				}
				result = append(result, content)
			}
		}
	}
	return
}
