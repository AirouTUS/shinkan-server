package output

import (
	"github.com/AirouTUS/shinkan-server/pkg/model"
	"github.com/AirouTUS/shinkan-server/pkg/usecase"
)

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

type Circle struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	About       string        `json:"about"`
	CatchCopy   string        `json:"catchCopy"`
	Description string        `json:"description"`
	EyeCatch    string        `json:"eyecatch"`
	Images      []CircleImage `json:"images"`
	Types       []CircleType  `json:"types"`
	Category    Category      `json:"category"`
}

type CircleImage struct {
	URL string `json:"url"`
}

type CircleType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToGetCircle(circles []*model.Circle, categories []*model.Category) (result Circle) {
	var circle model.Circle
	if len(circles) > 0 {
		circle = *circles[0]
	}

	result.ID = circle.ID
	result.Name = circle.Name
	result.About = circle.About
	result.CatchCopy = circle.CatchCopy
	result.Description = circle.Description
	result.EyeCatch = circle.EyeCatch

	for _, v := range categories {
		if v.ID == circle.CategoryID {
			result.Category.ID = v.ID
			result.Category.Name = v.Name
		}
	}

	result.Types = make([]CircleType, 0, len(circles))
	for _, v := range circles {
		content := CircleType{
			ID:   *v.TypeID,
			Name: *v.TypeName,
		}
		result.Types = append(result.Types, content)
	}

	result.Images = make([]CircleImage, 0)
	return
}

type CircleList struct {
	Circles []Circle `json:"circles"`
}

func ToListCircle(circles []*usecase.Circle) (result CircleList) {
	result.Circles = make([]Circle, 0, len(circles))
	for _, v := range circles {
		content := Circle{
			ID:          v.ID,
			Name:        v.Name,
			About:       v.About,
			CatchCopy:   v.CatchCopy,
			EyeCatch:    v.EyeCatch,
			Description: v.Description,
			Category:    Category(v.Category),
		}
		content.Types = make([]CircleType, 0, len(v.Types))
		for _, vv := range v.Types {
			content.Types = append(content.Types, CircleType(vv))
		}
		content.Images = make([]CircleImage, 0, len(v.Images))
		for _, vv := range v.Images {
			content.Images = append(content.Images, CircleImage(vv))
		}
		result.Circles = append(result.Circles, content)
	}
	return
}
