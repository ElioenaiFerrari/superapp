package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ElioenaiFerrari/superapp/internal/services"
	fiber "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	v1 := app.Group("/api/v1")

	restaurantsHost := os.Getenv("RESTAURANTS_HOST")
	rConn, err := grpc.Dial(restaurantsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer rConn.Close()
	rc := services.NewRestaurantServiceClient(rConn)

	catalogsHost := os.Getenv("CATALOGS_HOST")
	cConn, err := grpc.Dial(catalogsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer cConn.Close()
	cc := services.NewCatalogServiceClient(cConn)

	platesHost := os.Getenv("PLATES_HOST")
	pConn, err := grpc.Dial(platesHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer pConn.Close()

	pc := services.NewPlateServiceClient(pConn)

	ordersHost := os.Getenv("ORDERS_HOST")
	oConn, err := grpc.Dial(ordersHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer pConn.Close()

	oc := services.NewOrderServiceClient(oConn)

	v1.Get("/restaurants", func(c fiber.Ctx) error {
		ctx := c.Context()
		restaurants, err := rc.ListRestaurants(ctx, &services.Empty{})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(restaurants.Restaurants)
	})

	v1.Get("/restaurants/:id", func(c fiber.Ctx) error {
		ctx := c.Context()
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		restaurant, err := rc.GetRestaurant(ctx, &services.GetRestaurantRequest{Id: id})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(restaurant)
	})

	v1.Get("/restaurants/:id/catalogs", func(c fiber.Ctx) error {
		ctx := c.Context()
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		catalogs, err := cc.ListCatalogs(ctx, &services.ListCatalogsRequest{RestaurantId: id})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(catalogs.Catalogs)
	})

	v1.Get("/catalogs/:id/plates", func(c fiber.Ctx) error {
		ctx := c.Context()
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		plates, err := pc.ListPlates(ctx, &services.ListPlatesRequest{CatalogId: id})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(plates.Plates)
	})

	v1.Get("/catalogs/:id/qrcode", func(c fiber.Ctx) error {
		ctx := c.Context()
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		qrcode, err := cc.GetCatalogQRCode(ctx, &services.GetCatalogRequest{Id: id})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		c.Context().SetContentType("image/png")

		return c.Send(qrcode.QrCode)
	})

	type createOrderDTO struct {
		PlateIDs []int64 `json:"plate_ids"`
	}

	v1.Post("/users/:id/orders", func(c fiber.Ctx) error {
		ctx := c.Context()
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		var createOrderDTO createOrderDTO
		if err := c.AutoFormat(&createOrderDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		order, err := oc.CreateOrder(ctx, &services.CreateOrderRequest{UserId: id, PlateIds: createOrderDTO.PlateIDs})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(order)
	})

	v1.Get("/restaurants/:id/orders", func(c fiber.Ctx) error {
		ctx := c.Context()
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		orders, err := oc.ListOrders(ctx, &services.ListOrdersRequest{RestaurantId: id})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(orders.Orders)
	})

	v1.Get("/users/:id/orders", func(c fiber.Ctx) error {
		ctx := c.Context()

		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		orders, err := oc.ListOrders(ctx, &services.ListOrdersRequest{UserId: id})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(orders.Orders)
	})

	port := os.Getenv("HOST")
	if port == "" {
		port = "4000"
	}
	log.Fatal(app.Listen(port))
}
