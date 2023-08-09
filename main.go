package main

import (
	"github.com/gin-gonic/gin"
	"resp-pgsql/controllers/productcontroller"
	"resp-pgsql/models"
)

func main() {

	route := gin.Default()
	models.ConnectDatabase()

	route.GET("/api/products", productcontroller.Index)
	route.GET("/api/product/:id", productcontroller.Show)
	route.POST("/api/product", productcontroller.Create)
	route.PUT("/api/product/:id", productcontroller.Update)
	route.DELETE("/api/product", productcontroller.Delete)
	route.DELETE("/api/product/deletebyid/:id", productcontroller.DeleteById)

	route.Run("localhost:9292")
}
