package main

import (
	"cluster-api-management-dashboard/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.GinRoutes(router)

	err := router.Run("localhost:8091")
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}
