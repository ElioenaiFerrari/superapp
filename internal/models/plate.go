package models

type Plate struct {
	ID        int64    `json:"id,omitempty"`
	CatalogID int64    `json:"catalog_id,omitempty"`
	Catalog   *Catalog `json:"catalog,omitempty"`
	ImageURL  string   `json:"image_url,omitempty"`
	Name      string   `json:"name,omitempty"`
	Price     int64    `json:"price,omitempty"`
}
