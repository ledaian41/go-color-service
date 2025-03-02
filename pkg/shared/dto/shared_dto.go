package shared_dto

type HSL struct {
	H, S, L float64 // Hue (0-360), Saturation (0-100), Lightness (0-100)
}

type CP6Response struct {
	Main            string `json:"primary"`
	Tint            string `json:"primary-light"`
	Shade           string `json:"primary-dark"`
	Complementary   string `json:"complementary"`
	Analogous       string `json:"analogous"`
	AnalogousSecond string `json:"analogous-second"`
}
