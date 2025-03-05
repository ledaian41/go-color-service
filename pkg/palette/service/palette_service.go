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
			adjustHSL(hsl, 0, 10, 20),   // Tint (Primary Light)
			adjustHSL(hsl, 0, -10, -20), // Shade (Primary Dark)
			adjustHSL(hsl, 180, 0, 0),   // Complementary
		}
	case 8:
		return []string{
			mainColor,
			adjustHSL(hsl, 0, 10, 20),    // Tint (Primary Light)
			adjustHSL(hsl, 0, -10, -20),  // Shade (Primary Dark)
			adjustHSL(hsl, 180, 0, 0),    // Complementary
			adjustHSL(hsl, 180, 10, -20), // Complementary Light
			adjustHSL(hsl, 180, -10, 20), // Complementary Dark
			adjustHSL(hsl, 30, 0, 0),     // Analogous 1
			adjustHSL(hsl, -30, 0, 0),    // Analogous 2
		}
	default:
		return []string{
			mainColor,
			adjustHSL(hsl, 0, 10, 20),   // Tint (Primary Light)
			adjustHSL(hsl, 0, -10, -20), // Shade (Primary Dark)
			adjustHSL(hsl, 180, 0, 0),   // Complementary
			adjustHSL(hsl, 30, 0, 0),    // Analogous 1
			adjustHSL(hsl, -30, 0, 0),   // Analogous 2
		}
	}
}

func adjustHSL(hsl *shared_dto.HSL, hAdj, sAdj, lAdj float64) string {
	h := math.Mod(hsl.H+hAdj, 360)              // Hue giữ trong 0-360
	s := math.Min(100, math.Max(0, hsl.S+sAdj)) // Saturation 0-100
	l := math.Min(100, math.Max(0, hsl.L+lAdj)) // Lightness 0-100
	return shared_utils.HslToHex(h, s, l)
}
