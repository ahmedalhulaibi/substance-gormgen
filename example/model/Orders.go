package model

type Order struct {
	OrderID     int32 `gorm:"column:OrderID;primary_key;" `
	OrderNumber int32 `gorm:"column:OrderNumber;" `
	PersonID    int32 `gorm:"column:PersonID;" `
}
