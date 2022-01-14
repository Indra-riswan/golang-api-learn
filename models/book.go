package models

type Book struct {
	ID          uint   `gorm:"primaryKey; autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(200)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint   `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
