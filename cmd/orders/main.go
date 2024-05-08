package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ElioenaiFerrari/superapp/internal/services"
	"github.com/ElioenaiFerrari/superapp/pkg/logger"
	"github.com/rotisserie/eris"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	services.UnimplementedOrderServiceServer
	restaurantService services.RestaurantServiceClient
	platesService     services.PlateServiceClient
	catalogService    services.CatalogServiceClient
	orders            []*services.Order
}

func (h *Handler) ListOrders(
	ctx context.Context,
	req *services.ListOrdersRequest,
) (*services.ListOrdersResponse, error) {
	orders := []*services.Order{}
	for _, o := range h.orders {
		if o.RestaurantId == req.RestaurantId || req.UserId == o.UserId {
			orders = append(orders, o)
		}
	}

	return &services.ListOrdersResponse{
		Orders: orders,
	}, nil
}

func (h *Handler) GetOrder(ctx context.Context, req *services.GetOrderRequest) (*services.Order, error) {
	for _, o := range h.orders {
		if o.Id == req.Id {
			return o, nil
		}
	}

	return nil, eris.New("order not found")
}

func (h *Handler) CreateOrder(ctx context.Context, req *services.CreateOrderRequest) (*services.Order, error) {
	order := &services.Order{
		Id:       int64(len(h.orders) + 1),
		UserId:   req.UserId,
		PlateIds: req.PlateIds,
		Status:   services.OrderStatus_CREATED,
	}

	for _, plateId := range req.PlateIds {
		plate, err := h.platesService.GetPlate(ctx, &services.GetPlateRequest{Id: plateId})
		if err != nil {
			return nil, eris.Wrap(err, "failed to get plate")
		}

		catalog, err := h.catalogService.GetCatalog(ctx, &services.GetCatalogRequest{Id: plate.CatalogId})
		if err != nil {
			return nil, eris.Wrap(err, "failed to get catalog")
		}

		if order.RestaurantId == 0 {
			order.RestaurantId = catalog.RestaurantId
		}

		if order.RestaurantId != catalog.RestaurantId {
			return nil, eris.New("plates from different restaurants")
		}

		order.Price += plate.Price
	}

	h.orders = append(h.orders, order)

	return order, nil
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
	restaurantsHost := os.Getenv("RESTAURANTS_HOST")
	rConn, err := grpc.Dial(restaurantsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer rConn.Close()
	rc := services.NewRestaurantServiceClient(rConn)

	platesHost := os.Getenv("PLATES_HOST")
	pConn, err := grpc.Dial(platesHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer pConn.Close()

	pc := services.NewPlateServiceClient(pConn)

	catalogsHost := os.Getenv("CATALOGS_HOST")
	cConn, err := grpc.Dial(catalogsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer cConn.Close()
	cc := services.NewCatalogServiceClient(cConn)

	services.RegisterOrderServiceServer(s, &Handler{
		restaurantService: rc,
		platesService:     pc,
		catalogService:    cc,
		orders:            []*services.Order{},
	})

	logger.L.Info().Str("port", port).Msg("catalogs service started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
