package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := Run(context.Background()); err != nil {
		panic(err)
	}
}
func Run(ctx context.Context) error {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return err
	}

	employeeDAO, err := NewEmployeeDAO(ctx, client)
	if err != nil {
		return err
	}

	memoryStorage := NewMemoryStorage(employeeDAO)
	handler := NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()

	return err
}
