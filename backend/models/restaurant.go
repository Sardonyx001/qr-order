package models

type Restaurant struct {
	UUIDBaseModel
	Name       string     `json:"name" gorm:"not null"`
	Users      []*User    `json:"users" gorm:"many2many:user_restaurants;"`
	Items      []Item     `json:"items"`
	Orders     []Order    `json:"ordes"`
	Categories []Category `json:"categories"`
	Tables     []Table    `json:"tables"`
}
