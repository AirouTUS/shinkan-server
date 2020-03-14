package output

import (
	"github.com/AirouTUS/shinkan-server/internal/model"
	"github.com/AirouTUS/shinkan-server/internal/usecase"
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

type CircleList struct {
	Categories []Category `json:"categories"`
	Circles    []Circle   `json:"circles"`
}

func ToListCircle(circles []*model.Circle, categories []usecase.CircleCategory) (result CircleList) {
	result.Categories = make([]Category, 0, len(categories))
	result.Circles = make([]Circle, 0, len(circles))
	for _, v := range circles {
		content := Circle{
			ID:          v.ID,
			Name:        v.Name,
			About:       v.About,
			CatchCopy:   v.CatchCopy,
			Description: v.Description,
		}
		result.Circles = append(result.Circles, content)
	}
	for _, v := range categories {
		content := Category{
			ID:   v.ID,
			Name: v.Name,
		}
		result.Categories = append(result.Categories, content)
	}
	return
}
