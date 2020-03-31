package model

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
	CategoryID    int
	Email         string
	Twitter       string
	URL           string
	EyeCatch      string
	TypeID        *int
	TypeName      *string
}

type GetCircle struct {
	ID            int
	Name          string
	About         string
	CatchCopy     string
	Cost          *string
	Location      *string
	WorkTime      *string
	MembersNumber *string
	Description   string
	CategoryID    int
	Email         string
	Twitter       string
	URL           string
	EyeCatch      string
	Types         []CircleType
	Images        []CircleImages
}

type CircleType struct {
	ID   int
	Name string
}

type CircleImages struct {
	URL string
}
