package model

type Category struct {
	ID       uint      `gorm:"PrimaryKey" json:"id"`
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products"`
}
