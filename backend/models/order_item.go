package models

type State string

const (
	READY     State = "READY"
	COMING    State = "COMING"
	CANCELLED State = "CANCELLED"
)

type OrderItem struct {
	UUIDBaseModel
	Quantity uint   `json:"quantity" gorm:"default:1"`
	Status   State  `json:"status"   gorm:"default:'COMING';"`
	Options  string `json:"options"`
	OrderID  string `json:"order_id"`
	ItemID   string `json:"item_id"`
}
