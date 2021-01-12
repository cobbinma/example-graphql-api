package models

import "time"

type MenuItemInput struct {
	// ID is the unique ID of the Menu Item.
	ID string `json:"id"`
	// Status is the current availability of the item. Item is available if not listed.
	Status ItemStatus `json:"status"`
}

func (m MenuItemInput) MenuItem(now time.Time) MenuItem {
	return MenuItem{
		ID:          m.ID,
		Status:      m.Status,
		AvailableAt: m.availableAt(now),
	}
}

func (m MenuItemInput) availableAt(now time.Time) (availableAt *time.Time) {
	if m.Status == ItemStatusHidden {
		return nil
	}

	addTwoHours := now.Add(2 * time.Hour)
	midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	if addTwoHours.After(midnight) {
		return &midnight
	}

	return &addTwoHours
}
