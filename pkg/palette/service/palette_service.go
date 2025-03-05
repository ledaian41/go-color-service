package palette_service

import (
	"github.com/ledaian41/go-color-service/pkg/shared/dto"
	"github.com/ledaian41/go-color-service/pkg/shared/utils"
	"math"
)

type PaletteService struct{}

func NewPaletteService() *PaletteService {
	return &PaletteService{}
}

func (c PaletteService) GenerateColorPalette(mainColor string, version int8) []string {
	if valid := shared_utils.IsValidHexColor(mainColor); !valid {
		return []string{}
	}
	hsl := shared_utils.HexToHSL(mainColor)
	if hsl == nil {
		return []string{}
	}

	switch version {
	case 4:
		return []string{
			mainColor,
			adjustHSL(hsl, 0, -20, 15),  // Tint (Primary Light)
			adjustHSL(hsl, 0, -15, -15), // Shade (Primary Dark)
			adjustHSL(hsl, 180, -30, 0), // Complementary
		}
	case 8:
		compHsl := shared_dto.HSL{H: math.Mod(hsl.H+180, 360), S: hsl.S - 30, L: hsl.L}
		return []string{
			mainColor,
			adjustHSL(hsl, 0, -20, 15),      // Tint (Primary Light)
			adjustHSL(hsl, 0, -15, -15),     // Shade (Primary Dark)
			adjustHSL(&compHsl, 0, 0, 0),    // Complementary
			adjustHSL(&compHsl, 0, -10, 15), // Complementary Light
			adjustHSL(&compHsl, 0, -5, -15), // Complementary Dark
			adjustHSL(hsl, 20, -10, 0),      // Analogous 1
			adjustHSL(hsl, -20, -10, 5),     // Analogous 2
		}
	default:
		return []string{
			mainColor,
			adjustHSL(hsl, 0, -20, 15),  // Tint (Primary Light)
			adjustHSL(hsl, 0, -15, -15), // Shade (Primary Dark)
			adjustHSL(hsl, 180, -30, 0), // Complementary
			adjustHSL(hsl, 20, -10, 0),  // Analogous 1
			adjustHSL(hsl, -20, -10, 5), // Analogous 2
		}
	}
}

func adjustHSL(hsl *shared_dto.HSL, hAdj, sAdj, lAdj float64) string {
	h := math.Mod(hsl.H+hAdj, 360)              // Hue giữ trong 0-360
	s := math.Min(100, math.Max(0, hsl.S+sAdj)) // Saturation 0-100
	l := math.Min(100, math.Max(0, hsl.L+lAdj)) // Lightness 0-100
	return shared_utils.HslToHex(h, s, l)
}
