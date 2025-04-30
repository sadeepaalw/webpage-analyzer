package main

import (
	"github.com/gin-gonic/gin"
	"web-analyzer/routes"
	"web-analyzer/utils"
)

func main() {
	utils.InitLogger()
	//config.LoadConfig()
	//database.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		utils.Log.Fatal("Server failed to start:", err)
	}
}
