package shared_dto

type HSL struct {
	H, S, L float64 // Hue (0-360), Saturation (0-100), Lightness (0-100)
}

type RGB struct {
	R, G, B float64
}

type CPResponse struct {
	Main               string `json:"primary"`
	Tint               string `json:"primary-light"`
	Shade              string `json:"primary-dark"`
	Complementary      string `json:"complementary"`
	ComplementaryLight string `json:"complementary-light,omitempty"`
	ComplementaryDark  string `json:"complementary-dark,omitempty"`
	Analogous          string `json:"analogous,omitempty"`
	AnalogousSecond    string `json:"analogous-second,omitempty"`
}
