package usecase

import (
	"fmt"
	"strings"

	"github.com/AirouTUS/shinkan-server/internal/model"
)

type CircleList []Circle

type Circle struct {
	ID            int
	Name          string
	About         string
	CatchCopy     string
	Cost          *string
	Location      *string
	WorkTime      *string
	MembersNumber *string
	Description   string
	EyeCatch      string
	Twitter       string
	Email         string
	URL           string
	Images        []CircleImage
	Types         []CircleType
	Category      Category
	UpdatedAt     string
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

func ParseCircles(circles []*model.Circle, categories []*model.Category, q string) (result CircleList) {
	var circle Circle
	b := 0
	for _, v := range circles {
		if b != v.ID {
			if b != 0 {
				cc := circle
				result = append(result, cc)
			}
			circle = Circle{
				ID:            v.ID,
				Name:          v.Name,
				About:         v.About,
				CatchCopy:     v.CatchCopy,
				Cost:          v.Cost,
				Location:      v.Location,
				WorkTime:      v.WorkTime,
				MembersNumber: v.MembersNumber,
				Description:   v.Description,
				EyeCatch:      v.EyeCatch,
				Twitter:       v.Twitter,
				Email:         v.Email,
				URL:           v.URL,
				Category:      Category{ID: v.CategoryID},
				UpdatedAt:     v.UpdatedAt,
			}
			circle.Types = make([]CircleType, 0)
			if v.TypeID != nil {
				cType := CircleType{
					ID:   *v.TypeID,
					Name: *v.TypeName,
				}
				circle.Types = append(circle.Types, cType)
			}
			b = v.ID
		} else {
			content := CircleType{
				ID:   *v.TypeID,
				Name: *v.TypeName,
			}
			circle.Types = append(circle.Types, content)
		}
	}
	result = append(result, circle)

	for i, v := range result {
		for _, vv := range categories {
			if v.Category.ID == vv.ID {
				result[i].Category.Name = vv.Name
			}
		}
	}

	return result.search(q)
}

func (c CircleList) search(q string) (result CircleList) {
	query := strings.Split(q, " ")

	for _, v := range query {
		if v == "" {
			continue
		}
		unmatches := make([]int, 0)
		for i, vv := range c {
			flag := false
			if strings.Contains(vv.Name, v) {
				flag = true
			}
			if strings.Contains(vv.About, v) {
				flag = true
			}
			if strings.Contains(vv.Description, v) {
				flag = true
			}
			for _, vvv := range vv.Types {
				if strings.Contains(vvv.Name, v) {
					flag = true
				}
			}
			if !flag {
				unmatches = append(unmatches, i)
			}
		}

		for i, j := range unmatches {
			k := j - i
			c = append(c[:k], c[k+1:]...)
		}
		fmt.Println(unmatches)
	}
	return c
}
