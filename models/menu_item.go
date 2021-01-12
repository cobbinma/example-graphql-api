package models

import "time"

type MenuItem struct {
	// ID is the unique ID of the Menu Item.
	ID string `json:"id" db:"id"`
	// Status is the current availability of the item. Item is available if not listed.
	Status ItemStatus `json:"status" db:"status"`
	// Available at is the time the menu item will become available. Only valid for unavailable items.
	AvailableAt *time.Time `json:"availableAt" db:"available_at"`
}
