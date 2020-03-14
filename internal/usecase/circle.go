package usecase

import (
	"github.com/AirouTUS/shinkan-server/internal/model"
)

type Circle struct {
	ID          int
	Name        string
	About       string
	CatchCopy   string
	Description string
	Images      []CircleImage
	Types       []CircleType
	Category    Category
}

type CircleImage struct {
	URL string `json:"url"`
}

type CircleType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int
	Name string
}

func ParseCircles(circles []*model.Circle, categories []*model.Category) (result []*Circle) {
	var circle Circle
	b := 0
	for _, v := range circles {
		if b != v.ID {
			if b != 0 {
				cc := circle
				result = append(result, &cc)
			}
			circle = Circle{
				ID:          v.ID,
				Name:        v.Name,
				About:       v.About,
				CatchCopy:   v.CatchCopy,
				Description: v.Description,
				Category:    Category{ID: v.CategoryID},
			}
			circle.Types = make([]CircleType, 0)
			b = v.ID
		} else {
			content := CircleType{
				ID:   *v.TypeID,
				Name: *v.TypeName,
			}
			circle.Types = append(circle.Types, content)
		}
	}
	result = append(result, &circle)

	for i, v := range result {
		for _, vv := range categories {
			if v.Category.ID == vv.ID {
				result[i].Category.Name = vv.Name
			}
		}
	}
	return
}
