package color_service

import (
	"github.com/ledaian41/go-color-service/pkg/shared/utils"
)

type ColorService struct{}

func NewColorService() *ColorService {
	return &ColorService{}
}

func (c ColorService) TextColor(background string) string {
	if valid := shared_utils.IsValidHexColor(background); !valid {
		return "#000000"
	}

	hsl := shared_utils.HexToHSL(background)
	if hsl == nil || hsl.L >= 50 {
		return "#000000"
	}

	return "#FFFFFF"
}
