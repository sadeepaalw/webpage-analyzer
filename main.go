package main

import (
	"web-analyzer/routes"
	"web-analyzer/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()

	r := gin.Default()
	routes.SetupRoutes(r)

	// todo only for dev
	_ = r.SetTrustedProxies([]string{"0.0.0.0/0"})

	err := r.Run(":8080")
	if err != nil {
		utils.Log.Fatal("Server failed to start:", err)
	}
}
