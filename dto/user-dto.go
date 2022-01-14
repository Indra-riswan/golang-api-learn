package dto

type DTORigister struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=3"`
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" form:"password"`
}
type DTOUpdateUser struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

type DTOLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
