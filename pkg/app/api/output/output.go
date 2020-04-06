package output

import (
	"log"
	"strconv"

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
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	About         string        `json:"about"`
	CatchCopy     string        `json:"catchCopy"`
	MembersNumber string        `json:"members_number"`
	WorkTime      string        `json:"workTime"`
	Location      string        `json:"location"`
	Cost          string        `json:"cost"`
	Description   string        `json:"description"`
	EyeCatch      string        `json:"eyecatch"`
	Twitter       string        `json:"twitter"`
	Email         string        `json:"email"`
	URL           string        `json:"url"`
	Images        []CircleImage `json:"images"`
	Types         []CircleType  `json:"types"`
	Category      Category      `json:"category"`
	UpdatedAt     string        `json:"updated_at"`
}

type CircleImage struct {
	URL string `json:"url"`
}

type CircleType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToGetCircle(circle *model.GetCircle, categories []*model.Category) (result Circle) {
	result.ID = circle.ID
	result.Name = circle.Name
	result.About = circle.About
	result.CatchCopy = circle.CatchCopy
	result.Description = circle.Description
	result.EyeCatch = circle.EyeCatch
	result.Twitter = circle.Twitter
	result.Email = circle.Email
	result.URL = circle.URL
	result.UpdatedAt = circle.UpdatedAt

	if circle.MembersNumber != nil {
		result.MembersNumber = *circle.MembersNumber
	}
	if circle.WorkTime != nil {
		result.WorkTime = *circle.WorkTime
	}
	if circle.Location != nil {
		result.Location = *circle.Location
	}
	if circle.Cost != nil {
		result.Cost = *circle.Cost
	}

	for _, v := range categories {
		if v.ID == circle.CategoryID {
			result.Category.ID = v.ID
			result.Category.Name = v.Name
		}
	}

	result.Types = make([]CircleType, 0, len(circle.Types))
	for _, v := range circle.Types {
		content := CircleType{
			ID:   v.ID,
			Name: v.Name,
		}
		result.Types = append(result.Types, content)
	}

	result.Images = make([]CircleImage, 0, len(circle.Images))
	for _, v := range circle.Images {
		content := CircleImage{
			URL: v.URL,
		}
		result.Images = append(result.Images, content)
	}
	return
}

type CircleList struct {
	Circles []Circle `json:"circles"`
}

func ToListCircle(startStr, endStr string, circles usecase.CircleList) (result CircleList) {
	result.Circles = make([]Circle, 0, len(circles))

	var start int
	var sliceEnd int
	const perPage = 10
	start, err := strconv.Atoi(startStr)
	if err != nil {
		log.Println(err)
		return
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		log.Println(err)
		return
	}

	if start >= len(circles) {
		return
	}

	if end == 0 || end >= len(circles) {
		sliceEnd = len(circles)
	} else {
		sliceEnd = end + 1
	}
	if start == 0 && end == 0 {
		sliceEnd = 1
	}

	if startStr == "-1" && endStr == "-1" {
		start = 0
		sliceEnd = perPage
		if len(circles) < perPage {
			sliceEnd = len(circles)
		}
	}
	for _, v := range circles[start:sliceEnd] {
		content := Circle{
			ID:          v.ID,
			Name:        v.Name,
			About:       v.About,
			CatchCopy:   v.CatchCopy,
			EyeCatch:    v.EyeCatch,
			Twitter:     v.Twitter,
			Email:       v.Email,
			URL:         v.URL,
			Description: v.Description,
			Category:    Category(v.Category),
			UpdatedAt:   v.UpdatedAt,
		}
		if v.MembersNumber != nil {
			content.MembersNumber = *v.MembersNumber
		}
		if v.WorkTime != nil {
			content.WorkTime = *v.WorkTime
		}
		if v.Location != nil {
			content.Location = *v.Location
		}
		if v.Cost != nil {
			content.Cost = *v.Cost
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
