package models

type Catalog struct {
	ID           int64       `json:"id,omitempty"`
	RestaurantID int64       `json:"restaurant_id,omitempty"`
	Restaurant   *Restaurant `json:"restaurant,omitempty"`
	Plates       []*Plate    `json:"plates,omitempty"`
	Name         string      `json:"name,omitempty"`
}
