package model

type Person struct {
	Address   string    `gorm:"column:Address;" `
	AntiOrder AntiOrder `gorm:"ForeignKey:PersonID;AssociationForeignKey:ID;" `
	City      string    `gorm:"column:City;" `
	FirstName string    `gorm:"column:FirstName;" `
	ID        int32     `gorm:"column:ID;primary_key;" `
	LastName  string    `gorm:"column:LastName;" `
	Orders    []Order   `gorm:"ForeignKey:PersonID;AssociationForeignKey:ID;" `
}
