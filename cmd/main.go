package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/color/handler"
	"github.com/ledaian41/go-color-service/pkg/color/service"
	"github.com/ledaian41/go-color-service/pkg/palette/handler"
	"github.com/ledaian41/go-color-service/pkg/palette/service"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	paletteService := palette_service.NewPaletteService()
	paletteHandler := palette_handler.NewPaletteHandler(paletteService)

	r.GET("/palette/:v/generate", paletteHandler.GenerateColorPalette)
	r.GET("/palette/:v/random", paletteHandler.RandomPalette)

	colorService := color_service.NewColorService()
	colorHandler := color_handler.NewColorHandler(colorService)
	r.GET("/color/text", colorHandler.TextColor)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port 8080
	}
	r.Run("0.0.0.0:" + port)
}
