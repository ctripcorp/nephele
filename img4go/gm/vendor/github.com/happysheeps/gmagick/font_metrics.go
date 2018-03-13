package gmagick

/*
%    Element Description
%    -------------------------------------------------
%          0 character width
%          1 character height
%          2 ascender
%          3 descender
%          4 text width
%          5 text height
%          6 maximum horizontal advance
*/
type FontMetrics struct {
	CharacterWidth           float64
	CharacterHeight          float64
	Ascender                 float64
	Descender                float64
	TextWidth                float64
	TextHeight               float64
	MaximumHorizontalAdvance float64
}

func NewFontMetricsFromArray(arr []float64) *FontMetrics {
	if len(arr) != 7 {
		panic("Wrong number of font metric items")
	}
	return &FontMetrics{arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6]}
}
