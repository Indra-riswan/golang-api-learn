package models

type User struct {
	ID       uint   `gorm:"primaryKey; autoIncrement" json:"id" form:"id" binding:"required"`
	Name     string `gorm:"type:char(50)" json:"name" form:"name" binding:"required"`
	Email    string `gorm:"type:varchar(100);unique" json:"email" form:"name" binding:"required"`
	Password string `gorm:"->;<-; not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
