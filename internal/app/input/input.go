package input

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type ListCircleInput struct {
	CategoryID []int `json:"categoryID"`
}

func (i ListCircleInput) Validate() error {
	for _, v := range i.CategoryID {
		if v <= 0 {
			return ErrInvalidRequest
		}
	}
	return nil
}
