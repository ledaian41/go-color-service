package color_service

import (
	"github.com/ledaian41/go-color-service/pkg/shared/dto"
	"github.com/ledaian41/go-color-service/pkg/shared/utils"
	"math"
)

type ColorService struct{}

func NewColorService() *ColorService {
	return &ColorService{}
}

func (c ColorService) TextColor(background string) string {
	if valid := shared_utils.IsValidHexColor(background); !valid {
		return "#000000"
	}

	rgb := shared_utils.HexToRgb(background)
	if rgb == nil {
		return "#000000"
	}

	if getLuminance(rgb) > 0.5 {
		return "#000000"
	}

	return "#FFFFFF"
}

func getLuminance(rgb *shared_dto.RGB) float64 {
	var rLin, gLin, bLin float64
	if rgb.R <= 0.03928 {
		rLin = rgb.R / 12.92
	} else {
		rLin = math.Pow((rgb.R+0.055)/1.055, 2.4)
	}
	if rgb.G <= 0.03928 {
		gLin = rgb.G / 12.92
	} else {
		gLin = math.Pow((rgb.G+0.055)/1.055, 2.4)
	}
	if rgb.B <= 0.03928 {
		bLin = rgb.B / 12.92
	} else {
		bLin = math.Pow((rgb.B+0.055)/1.055, 2.4)
	}

	return 0.2126*rLin + 0.7152*gLin + 0.0722*bLin
}
