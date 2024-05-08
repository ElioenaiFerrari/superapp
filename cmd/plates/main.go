package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ElioenaiFerrari/superapp/internal/services"
	"github.com/ElioenaiFerrari/superapp/pkg/data"
	"github.com/ElioenaiFerrari/superapp/pkg/logger"
	"google.golang.org/grpc"
)

type Handler struct {
	services.UnimplementedCatalogServiceServer
}

func (h *Handler) ListPlates(ctx context.Context, req *services.ListPlatesRequest) (*services.ListPlatesResponse, error) {
	var plates []*services.Plate
	for _, p := range data.Plates {
		if p.CatalogID == req.CatalogId {
			plates = append(plates, &services.Plate{
				Id:        p.ID,
				CatalogId: p.CatalogID,
				Name:      p.Name,
				Price:     p.Price,
				ImageUrl:  p.ImageURL,
			})
		}
	}

	return &services.ListPlatesResponse{
		Plates: plates,
	}, nil
}

func (h *Handler) GetPlate(ctx context.Context, req *services.GetPlateRequest) (*services.Plate, error) {
	for _, p := range data.Plates {
		if p.ID == req.Id {
			return &services.Plate{
				Id:        p.ID,
				CatalogId: p.CatalogID,
				Name:      p.Name,
				Price:     p.Price,
				ImageUrl:  p.ImageURL,
			}, nil
		}
	}

	return nil, nil
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
	services.RegisterPlateServiceServer(s, &Handler{})

	logger.L.Info().Str("port", port).Msg("plates service started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
