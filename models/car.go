package models

import (
	"mime/multipart"
)

type CarInput struct {
	Name string `form:"name"`
	CarType string `form:"carType"`
	Rating float32 `form:"rating" sql:"type:decimal(2,1)"`
	Fuel string `form:"fuel"`
	Image *multipart.FileHeader `form:"image"`
	HourRate int64 `form:"hourRate"`
	DayRate int64 `form:"dayRate"`
	MonthRate int64 `form:"monthRate"`
}

type Car struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"`
	Rating float32 `json:"rating" sql:"type:decimal(2,1)"`
	Fuel string `json:"fuel"`
	Image string `json:"image"`
	HourRate int64 `json:"hourRate"`
	DayRate int64 `json:"dayRate"`
	MonthRate int64 `json:"monthRate"`
}