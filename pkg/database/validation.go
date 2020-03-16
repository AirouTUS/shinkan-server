package database

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
)

func (i ListCategoryInput) validate() error {
	return nil
}

func (i GetCircleInput) validate() error {
	if i.ID <= 0 {
		return ErrInvalidInput
	}
	return nil
}

func (i ListCircleInput) validate() error {
	for _, v := range i.CategoryID {
		if v <= 0 {
			return ErrInvalidInput
		}
	}
	return nil
}
