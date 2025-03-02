package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/color/handler"
	"github.com/ledaian41/go-color-service/pkg/color/service"
)

func main() {
	r := gin.Default()

	colorService := color_service.NewColorService()
	colorHandler := color_handler.NewColorHandler(colorService)

	r.GET("/palette/generate", colorHandler.GenerateColorPalette)
	r.GET("/palette/random", colorHandler.RandomPalette)

	r.Run("localhost:8080")
}
