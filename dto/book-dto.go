package dto

type DTOBookregister struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint   `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type DTOBookUpdate struct {
	ID          uint   `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint   `json:"user_id,omitempty" form:"user_id,mitempty"`
}
