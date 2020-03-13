package output

import (
	"github.com/AirouTUS/shinkan-server/internal/model"
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
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	About       string       `json:"about"`
	CatchCopy   string       `json:"catchCopy"`
	Description string       `json:"description"`
	Types       []CircleType `json:"types"`
	Category    Category     `json:"category"`
}

type CircleType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToGetCircle(circle *model.Circle, types []*model.CirclesCircleTypes, categories []*model.Category) (result Circle) {
	result.ID = circle.ID
	result.Name = circle.Name
	result.About = circle.About
	result.CatchCopy = circle.CatchCopy
	result.Description = circle.Description

	for _, v := range categories {
		if v.ID == circle.CategoryID {
			result.Category.ID = v.ID
			result.Category.Name = v.Name
		}
	}
	for _, v := range types {
		content := CircleType{
			ID:   v.CircleTypeID,
			Name: v.Name,
		}
		result.Types = append(result.Types, content)
	}
	return
}
