package color_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/shared/interface"
	"github.com/ledaian41/go-color-service/pkg/shared/utils"
	"net/http"
)

type ColorHandler struct {
	colorService shared_interface.ColorServiceInterface
}

func NewColorHandler(colorService shared_interface.ColorServiceInterface) *ColorHandler {
	return &ColorHandler{colorService}
}

func (h *ColorHandler) RandomPalette(c *gin.Context) {
	randomColor := shared_utils.RandomHexColor()
	colorPalette := h.colorService.GenerateColorPalette(randomColor)
	c.JSON(http.StatusOK, shared_utils.ToColorPaletteResponse(colorPalette))
}

func (h *ColorHandler) GenerateColorPalette(c *gin.Context) {
	baseColor := c.Query("base")
	colorPalette := h.colorService.GenerateColorPalette(baseColor)
	if len(colorPalette) == 6 {
		c.JSON(http.StatusOK, shared_utils.ToColorPaletteResponse(colorPalette))
		return
	}
	c.JSON(http.StatusBadRequest, fmt.Sprintf("🎉 Sorry, Could not generate color palette!"))
}
