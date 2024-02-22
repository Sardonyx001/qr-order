package models

type Restaurant struct {
	UUIDBaseModel
	UserID     string
	User       User
	Items      []Item
	Orders     []Order
	Categories []Category
	Tables     []Table
}
