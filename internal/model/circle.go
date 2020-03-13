package model

type Circle struct {
	ID          int
	Name        string
	About       string
	CatchCopy   string
	Description string
	CategoryID  int
	Email       string
	Twitter     string
	URL         string
}

type CirclesCircleTypes struct {
	CircleTypeID int
	Name         string
}
