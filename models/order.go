package models

type OrderInput struct {
	CarId           int64  `json:"carId"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
	PickupDate      string `json:"pickupDate"`
	DropoffDate     string `json:"dropoffDate"`
	PickupTime      string `json:"pickupTime"`
}

type Order struct {
	Id              int64  `json:"id" gorm:"primaryKey"`
	CarId           int64  `json:"carId"`
	UserId          int64  `json:"userId"`
	AdminId         int64  `json:"adminId"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
	PickupDate      string `json:"pickupDate"`
	DropoffDate     string `json:"dropoffDate"`
	PickupTime      string `json:"pickupTime"`
}