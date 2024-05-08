package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ElioenaiFerrari/superapp/internal/services"
	"github.com/ElioenaiFerrari/superapp/pkg/data"
	"github.com/ElioenaiFerrari/superapp/pkg/logger"
	"github.com/skip2/go-qrcode"
	"google.golang.org/grpc"
)

type Handler struct {
	services.UnimplementedCatalogServiceServer
}

func (h *Handler) ListCatalogs(ctx context.Context, req *services.ListCatalogsRequest) (*services.ListCatalogsResponse, error) {
	var catalogs []*services.Catalog
	for _, c := range data.Catalogs {
		if c.RestaurantID == req.RestaurantId {
			catalogs = append(catalogs, &services.Catalog{
				Id:           c.ID,
				RestaurantId: c.RestaurantID,
				Name:         c.Name,
			})
		}
	}

	return &services.ListCatalogsResponse{
		Catalogs: catalogs,
	}, nil
}

func (h *Handler) GetCatalog(ctx context.Context, req *services.GetCatalogRequest) (*services.Catalog, error) {
	for _, c := range data.Catalogs {
		if c.ID == req.Id {
			return &services.Catalog{
				Id:           c.ID,
				RestaurantId: c.RestaurantID,
				Name:         c.Name,
			}, nil
		}
	}

	return nil, nil
}

func (h *Handler) GetCatalogQRCode(ctx context.Context, req *services.GetCatalogRequest) (*services.GetCatalogQRCodeResponse, error) {
	url := fmt.Sprintf("https://1db5-2804-56c-a587-6c00-fc90-a83b-a0bd-64d5.ngrok-free.app/api/v1/catalogs/%d", req.Id)
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Fatalf("failed to generate QR code: %v", err)
	}

	for _, c := range data.Catalogs {
		if c.ID == req.Id {
			return &services.GetCatalogQRCodeResponse{
				QrCode: png,
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
	services.RegisterCatalogServiceServer(s, &Handler{})

	logger.L.Info().Str("port", port).Msg("catalogs service started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
