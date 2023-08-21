package model

type User struct {
	ID       uint      `gorm:"PrimaryKey" json:"id"`
	Username string    `gorm:"index" json:"username"`
	Password string    `json:"password"`
	Cart     []Product `gorm:"many2many:user_products" json:"cart"`
}
