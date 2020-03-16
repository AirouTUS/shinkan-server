package database

import "github.com/AirouTUS/shinkan-server/pkg/model"

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

type CircleList []Circle

func (m CircleList) circles() []*model.Circle {
	result := make([]*model.Circle, 0, len(m))
	for _, v := range m {
		content := model.Circle(v)
		result = append(result, &content)
	}
	return result
}

type Circle struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	About       string  `db:"about"`
	CatchCopy   string  `db:"catch_copy"`
	Description string  `db:"description"`
	CategoryID  int     `db:"circle_category_id"`
	Email       string  `db:"email"`
	Twitter     string  `db:"twitter"`
	URL         string  `db:"url"`
	TypeID      *int    `db:"type_id"`
	TypeName    *string `db:"type_name"`
}

func (m Circle) circle() *model.Circle {
	var result model.Circle
	result = model.Circle(m)
	return &result
}
