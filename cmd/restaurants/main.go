package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ElioenaiFerrari/superapp/internal/services"
	"github.com/ElioenaiFerrari/superapp/pkg/data"
	"github.com/ElioenaiFerrari/superapp/pkg/logger"

	"github.com/rotisserie/eris"

	"google.golang.org/grpc"
)

type Handler struct {
	services.UnimplementedRestaurantServiceServer
}

func (h *Handler) GetRestaurant(ctx context.Context, req *services.GetRestaurantRequest) (*services.Restaurant, error) {
	logger.L.Info().Int64("id", req.Id).Msg("GetRestaurant")
	for _, r := range data.Restaurants {
		if r.ID == req.Id {
			return &services.Restaurant{
				Id:       r.ID,
				Name:     r.Name,
				ImageUrl: r.ImageURL,
			}, nil
		}
	}

	logger.L.Error().Int64("id", req.Id).Msg("restaurant not found")

	return nil, eris.New("restaurant not found")
}

func (h *Handler) ListRestaurants(ctx context.Context, req *services.Empty) (*services.ListRestaurantsResponse, error) {
	logger.L.Info().Interface("req", req).Msg("ListRestaurants")
	var restaurants []*services.Restaurant
	for _, r := range data.Restaurants {
		restaurants = append(restaurants, &services.Restaurant{
			Id:       r.ID,
			Name:     r.Name,
			ImageUrl: r.ImageURL,
		})
	}

	return &services.ListRestaurantsResponse{
		Restaurants: restaurants,
	}, nil
}

func main() {
	port := os.Getenv("HOST")
	if port == "" {
		port = "4000"
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	services.RegisterRestaurantServiceServer(s, &Handler{})

	logger.L.Info().Str("port", port).Msg("restaurant service started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
