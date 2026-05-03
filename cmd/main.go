package main

import (
	"github.com/alisonferreira/APIGO/controller"
	"github.com/alisonferreira/APIGO/db"
	"github.com/alisonferreira/APIGO/repository"
	"github.com/alisonferreira/APIGO/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRrepository := repository.NewProductRepository(dbConnection)
	//camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRrepository)
	//camada de controller
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductsById)
	server.Run(":8000")

}
