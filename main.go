package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"endpointsService/handlers"
	"endpointsService/repositories"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()
	endpointsGroup := e.Group("/api/v1/endpoints")

	ctx := context.Background()
	firestoreClient := createClient(ctx, e)

	dbRepository := new(repositories.EndpointRepository)
	dbRepository.Client = firestoreClient

	endpointsGroup.GET("/", func(c echo.Context) error {
		return handlers.GetEndpoints(c, dbRepository)
	})

	endpointsGroup.GET("/:id/tree", func(c echo.Context) error {
		return handlers.GetEndpointsTree(c, dbRepository)
	})

	endpointsGroup.POST("/", func(c echo.Context) error {
		return handlers.CreateEndpoint(c, dbRepository)
	})

	endpointsGroup.PUT("/:id", func(c echo.Context) error {
		return handlers.UpdateEndpoint(c, dbRepository)
	})

	endpointsGroup.DELETE("/:id", func(c echo.Context) error {
		return handlers.DeleteEndpoint(c, dbRepository)
	})

	e.Logger.Fatal(e.Start(":5001"))

	defer func(db *firestore.Client) {
		err := db.Close()
		if err != nil {
			e.Logger.Fatal("Couldn't close Firestore client")
		}
	}(firestoreClient)
}

func createClient(ctx context.Context, e *echo.Echo) *firestore.Client {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		e.Logger.Fatalf("Failed to create client: %v", err)
	}

	return client
}
