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

func (c ColorService) GenerateColorPalette(mainColor string) []string {
	hsl := shared_utils.HexToHSL(mainColor)
	if hsl == nil {
		return []string{}
	}

	return []string{
		mainColor,
		adjustHSL(hsl, 0, 10, -20), // Tint
		adjustHSL(hsl, 0, -10, 20), // Shade
		adjustHSL(hsl, 180, 0, 0),  // Complementary
		adjustHSL(hsl, 30, 0, 0),   // Analogous 1
		adjustHSL(hsl, -30, 0, 0),  // Analogous 2
	}
}

func adjustHSL(hsl *shared_dto.HSL, hAdj, sAdj, lAdj float64) string {
	h := math.Mod(hsl.H+hAdj, 360)              // Hue giữ trong 0-360
	s := math.Min(100, math.Max(0, hsl.S+sAdj)) // Saturation 0-100
	l := math.Min(100, math.Max(0, hsl.L+lAdj)) // Lightness 0-100
	return shared_utils.HslToHex(h, s, l)
}
