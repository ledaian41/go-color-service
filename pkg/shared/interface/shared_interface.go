package shared_interface

type PaletteServiceInterface interface {
	GenerateColorPalette(string, int8) []string
}
