package dto

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,alphaspace,min=2"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"` // yyyy-mm-dd
}
