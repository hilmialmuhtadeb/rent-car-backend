package models

import (
	"mime/multipart"
	"github.com/shopspring/decimal"
)

type CarInput struct {
	Name string `form:"name"`
	CarType string `form:"carType"`
	Rating decimal.Decimal `form:"rating" sql:"type:decimal(2,1)"`
	Fuel string `form:"fuel"`
	Image *multipart.FileHeader `form:"image"`
	HourRate int64 `form:"hourRate"`
	DayRate int64 `form:"dayRate"`
	MonthRate int64 `form:"monthRate"`
}