package color_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ledaian41/go-color-service/pkg/shared/interface"
	"net/http"
)

type ColorHandler struct {
	colorService shared_interface.ColorServiceInterface
}

func NewColorHandler(colorService shared_interface.ColorServiceInterface) *ColorHandler {
	return &ColorHandler{colorService}
}

func (h *ColorHandler) TextColor(c *gin.Context) {
	background := c.Query("background")

	textColor := h.colorService.TextColor(background)
	c.JSON(http.StatusOK, textColor)
}
