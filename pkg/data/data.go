package data

import "github.com/ElioenaiFerrari/superapp/internal/models"

var (
	Restaurants = []*models.Restaurant{
		{
			ID:       1,
			Name:     "Restaurant 1",
			ImageURL: "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8cmVzdGF1cmFudHxlbnwwfHwwfHx8MA%3D%3D",
		},
		{
			ID:       2,
			Name:     "Restaurant 2",
			ImageURL: "https://images.unsplash.com/photo-1590846406792-0adc7f938f1d?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OHx8cmVzdGF1cmFudHxlbnwwfHwwfHx8MA%3D%3D",
		},
	}

	Catalogs = []*models.Catalog{
		{
			ID:           1,
			RestaurantID: 1,
			Name:         "Catalog 1",
		},
		{
			ID:           2,
			RestaurantID: 1,
			Name:         "Catalog 2",
		},
		{
			ID:           3,
			RestaurantID: 2,
			Name:         "Catalog 3",
		},
	}

	Plates = []*models.Plate{
		{
			ID:        1,
			CatalogID: 1,
			Name:      "Yakisoba",
			Price:     1000,
			ImageURL:  "https://images.unsplash.com/photo-1612927601601-6638404737ce?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8Zm9vZCUyMHBsYXRlfGVufDB8fDB8fHww",
		},
		{
			ID:        2,
			CatalogID: 1,
			Name:      "Mistura",
			Price:     2000,
			ImageURL:  "https://images.unsplash.com/photo-1540432797114-187727adf19b?q=80&w=3174&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		},
		{
			ID:        3,
			CatalogID: 2,
			Name:      "Peito de frango artesanal",
			Price:     2000,
			ImageURL:  "https://images.unsplash.com/photo-1514326640560-7d063ef2aed5?q=80&w=2963&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		},
	}
)
