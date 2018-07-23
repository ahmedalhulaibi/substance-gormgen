package model

type AntiOrder struct {
	AntiOrderID     int32 `gorm:"column:AntiOrderID;primary_key;" `
	AntiOrderNumber int32 `gorm:"column:AntiOrderNumber;" `
	PersonID        int32 `gorm:"column:PersonID;" `
}
