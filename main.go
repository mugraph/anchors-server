package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/janebuoy/anchors-server/controllers"
	"github.com/janebuoy/anchors-server/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/scenes", controllers.FindScenes)
	r.POST("/scenes", controllers.CreateScene)

	r.GET("/properties", controllers.FindProperties)

	r.Run()
}
