package output

import "github.com/AirouTUS/shinkan-server/internal/model"

type CategoryList struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryList(categories []*model.Category) (result CategoryList) {
	result.Categories = make([]Category, 0, len(categories))
	for _, v := range categories {
		content := Category{
			ID:   v.ID,
			Name: v.Name,
		}
		result.Categories = append(result.Categories, content)
	}
	return
}
