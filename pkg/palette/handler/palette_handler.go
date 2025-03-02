package palette_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/shared/interface"
	"github.com/ledaian41/go-color-service/pkg/shared/utils"
	"net/http"
)

type PaletteHandler struct {
	paletteService shared_interface.PaletteServiceInterface
}

func NewPaletteHandler(paletteService shared_interface.PaletteServiceInterface) *PaletteHandler {
	return &PaletteHandler{paletteService}
}

func (h *PaletteHandler) RandomPalette(c *gin.Context) {
	randomColor := shared_utils.RandomHexColor()
	colorPalette := h.paletteService.GenerateColorPalette(randomColor)
	c.JSON(http.StatusOK, shared_utils.ToColorPaletteResponse(colorPalette))
}

func (h *PaletteHandler) GenerateColorPalette(c *gin.Context) {
	baseColor := c.Query("base")
	colorPalette := h.paletteService.GenerateColorPalette(baseColor)
	if len(colorPalette) == 6 {
		c.JSON(http.StatusOK, shared_utils.ToColorPaletteResponse(colorPalette))
		return
	}
	c.JSON(http.StatusBadRequest, fmt.Sprintf("🎉 Sorry, Could not generate color palette!"))
}
