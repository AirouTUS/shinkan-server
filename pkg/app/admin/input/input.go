package input

import (
	"errors"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type circleImage struct {
	URL string `json:"url"`
}

type circleType struct {
	ID int `json:"id"`
}

type PostCircleInput struct {
	Name        string        `json:"name"`
	About       string        `json:"about"`
	CatchCopy   string        `json:"catchCopy"`
	Description string        `json:"description"`
	EyeCatch    string        `json:"eyecatch"`
	Email       string        `json:"email"`
	Twitter     string        `json:"twitter"`
	URL         string        `json:"url"`
	Images      []circleImage `json:"images"`
	Types       []circleType  `json:"types"`
	CategoryID  int           `json:"category_id"`
}

func (i PostCircleInput) Validate() error {
	if i.Name == "" {
		return ErrInvalidRequest
	}
	if i.CategoryID <= 0 {
		return ErrInvalidRequest
	}
	return nil
}
