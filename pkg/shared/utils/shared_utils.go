package shared_utils

import (
	"fmt"
	"github.com/ledaian41/go-color-service/pkg/shared/dto"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func RandomHexColor() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func HexToHSL(hex string) *shared_dto.HSL {
	r, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		return nil
	}

	g, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		return nil
	}

	b, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		return nil
	}

	rFloat := float64(r) / 255
	gFloat := float64(g) / 255
	bFloat := float64(b) / 255

	maxV := math.Max(rFloat, math.Max(gFloat, bFloat))
	minV := math.Min(rFloat, math.Min(gFloat, bFloat))
	l := (maxV + minV) / 2 // Lightness

	var h, s float64
	if maxV == minV {
		h = 0
		s = 0
	} else {
		d := maxV - minV
		// Saturation
		if l > 0.5 {
			s = d / (2 - maxV - minV)
		} else {
			s = d / (maxV + minV)
		}

		// Hue
		switch maxV {
		case rFloat:
			h = (gFloat - bFloat) / d
			if gFloat < bFloat {
				h += 6
			}
		case gFloat:
			h = (bFloat-rFloat)/d + 2
		case bFloat:
			h = (rFloat-gFloat)/d + 4
		}
		h /= 6
	}

	result := shared_dto.HSL{H: h * 360, S: s * 100, L: l * 100}
	return &result
}

func HslToHex(h, s, l float64) string {
	s /= 100
	l /= 100
	c := (1 - math.Abs(2*l-1)) * s // Chroma
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r, g, b float64
	switch {
	case 0 <= h && h < 60:
		r, g, b = c, x, 0
	case 60 <= h && h < 120:
		r, g, b = x, c, 0
	case 120 <= h && h < 180:
		r, g, b = 0, c, x
	case 180 <= h && h < 240:
		r, g, b = 0, x, c
	case 240 <= h && h < 300:
		r, g, b = x, 0, c
	case 300 <= h && h < 360:
		r, g, b = c, 0, x
	}

	rInt := int(math.Round((r + m) * 255))
	gInt := int(math.Round((g + m) * 255))
	bInt := int(math.Round((b + m) * 255))
	return fmt.Sprintf("#%02X%02X%02X", rInt, gInt, bInt)
}

func ToColorPaletteResponse(colors []string) shared_dto.CP6Response {
	return shared_dto.CP6Response{
		Main:            colors[0],
		Tint:            colors[1],
		Shade:           colors[2],
		Complementary:   colors[3],
		Analogous:       colors[4],
		AnalogousSecond: colors[5],
	}
}
