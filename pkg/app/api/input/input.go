package input

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type ListCircleInput struct {
	Start      int   `json:"start"`
	End        int   `json:"end"`
	CategoryID []int `json:"categoryID"`
}

func (i ListCircleInput) Validate() error {
	if i.Start < 0 {
		return ErrInvalidRequest
	}
	if i.End < 0 {
		return ErrInvalidRequest
	}
	for _, v := range i.CategoryID {
		if v <= 0 {
			return ErrInvalidRequest
		}
	}
	return nil
}
