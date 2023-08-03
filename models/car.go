package models

import (
	"github.com/shopspring/decimal"
)

type Car struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"`
	Rating decimal.Decimal `json:"rating" sql:"type:decimal(2,1)"`
	Fuel string `json:"fuel"`
	Image string `json:"image"`
	HourRate int64 `json:"hourRate"`
	DayRate int64 `json:"dayRate"`
	MonthRate int64 `json:"monthRate"`
}