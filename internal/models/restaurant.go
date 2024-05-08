package models

type Restaurant struct {
	ID       int64      `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	ImageURL string     `json:"image_url,omitempty"`
	Catalogs []*Catalog `json:"catalogs,omitempty"`
}
