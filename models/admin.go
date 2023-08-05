package models

type Admin struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
	Password string `json:"password"`
}