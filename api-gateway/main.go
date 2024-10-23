package main

import (
	"github.com/JerryLegend254/order-processing-system/api-gateway/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")

}
