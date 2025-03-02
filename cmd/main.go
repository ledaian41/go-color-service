package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/palette/handler"
	"github.com/ledaian41/go-color-service/pkg/palette/service"
)

func main() {
	r := gin.Default()

	paletteService := palette_service.NewPaletteService()
	paletteHandler := palette_handler.NewPaletteHandler(paletteService)

	r.GET("/palette/:v/generate", paletteHandler.GenerateColorPalette)
	r.GET("/palette/:v/random", paletteHandler.RandomPalette)

	r.Run("localhost:8080")
}
