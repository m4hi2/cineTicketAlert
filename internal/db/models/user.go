package models

type Role string

type User struct {
	DefaultModel
	Name     string `json:"name" gorm:"column:name;size:40"`
	Email    string `json:"email" gorm:"column:email;size:64;unique;index:idx_email"`
	Password string `json:"password" gorm:"column:password;"`
	Role     Role   `json:"role" gorm:"column:role;size:16"`
}
