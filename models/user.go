package models

type User struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	City string `json:"city"`
	Zip string `json:"zip"`
	Address string `json:"address"`
	Role int `json:"role"`
}