package palette_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/shared/interface"
	"github.com/ledaian41/go-color-service/pkg/shared/utils"
	"net/http"
	"slices"
	"strconv"
)

type PaletteHandler struct {
	paletteService shared_interface.PaletteServiceInterface
}

func NewPaletteHandler(paletteService shared_interface.PaletteServiceInterface) *PaletteHandler {
	return &PaletteHandler{paletteService}
}

func (h *PaletteHandler) RandomPalette(c *gin.Context) {
	v, err := strconv.Atoi(c.Param("v"))
	if err != nil || !slices.Contains([]int{4, 6, 8}, v) {
		v = 6
	}

	version := int8(v)
	randomColor := shared_utils.RandomHexColor()
	colorPalette := h.paletteService.GenerateColorPalette(randomColor, version)
	switch version {
	case 4:
		c.JSON(http.StatusOK, shared_utils.ToColorPalette4Response(colorPalette))
		break
	case 8:
		c.JSON(http.StatusOK, shared_utils.ToColorPalette8Response(colorPalette))
		break
	default:
		c.JSON(http.StatusOK, shared_utils.ToColorPalette6Response(colorPalette))
		break
	}
}

func (h *PaletteHandler) GenerateColorPalette(c *gin.Context) {
	baseColor := c.Query("base")
	v, err := strconv.Atoi(c.Param("v"))
	if err != nil || !slices.Contains([]int{4, 6, 8}, v) {
		v = 6
	}

	version := int8(v)
	colorPalette := h.paletteService.GenerateColorPalette(baseColor, version)
	if len(colorPalette) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Sorry, Could not generate color palette!"))
		return
	}

	switch version {
	case 4:
		c.JSON(http.StatusOK, shared_utils.ToColorPalette4Response(colorPalette))
		break
	case 8:
		c.JSON(http.StatusOK, shared_utils.ToColorPalette8Response(colorPalette))
		break
	default:
		c.JSON(http.StatusOK, shared_utils.ToColorPalette6Response(colorPalette))
		break
	}
}
