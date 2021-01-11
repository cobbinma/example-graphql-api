package models

type MenuItemInput struct {
	// ID is the unique ID of the Menu Item.
	ID string `json:"id"`
	// Status is the current availability of the item. Item is available if not listed.
	Status ItemStatus `json:"status"`
}
