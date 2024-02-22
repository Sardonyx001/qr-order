package models

type State string

const (
	Empty    State = "EMPTY"
	Occupied State = "OCCUPIED"
)

type Table struct {
	UUIDBaseModel
	RestaurantID string
	Restaurant   Restaurant
	Name         string `json:"name" gorm:"not null"`
	State        State  `json:"state" gorm:"not null"`
	Orders       []Order
}
