package stores

import "gorm.io/gorm"

type Stores struct {
	DB         *gorm.DB
	User       UserStore
	Admin      AdminStore
	Restaurant RestaurantStore
	Item       ItemStore
	Category   CategoryStore
}

func New(db *gorm.DB) *Stores {
	return &Stores{
		DB:         db,
		User:       &userStore{db},
		Admin:      &adminStore{db},
		Restaurant: &restaurantStore{db},
		Item:       &itemStore{db},
		Category:   &categoryStore{db},
	}
}
