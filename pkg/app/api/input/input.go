package input

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type ListCircleInput struct {
	Start      string `query:"start"`
	End        string `query:"end"`
	CategoryID []int  `query:"categoryID"`
	Q          string `query:"q"`
}

func (i ListCircleInput) Validate() error {
	start, err := strconv.Atoi(i.Start)
	if err != nil {
		return ErrInvalidRequest
	}
	if start < -1 {
		return ErrInvalidRequest
	}
	end, err := strconv.Atoi(i.End)
	if err != nil {
		return ErrInvalidRequest
	}
	if end < -1 {
		return ErrInvalidRequest
	}
	for _, v := range i.CategoryID {
		if v <= 0 {
			return ErrInvalidRequest
		}
	}
	return nil
}
