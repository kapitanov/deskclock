package main

func InitializeWeather() error {
	return nil
}

func DrawWeather(r Renderer, f Font, icon Font) {
	const (
		margin = 1
	)

	// TODO placeholder for future weather display
	s := "CLOCK"
	w, h := f.Measure(s)
	x := (r.Width() - w) / 2
	y := r.Height() - h - margin

	f.Render(r, x, y, s)

	return

	// icons:
	// 0 - rain
	// 1 - cloud
	// 2 - sun
	// 3 - storm
	// 4 - snow
	// 5 - heavy rain

	// str := "+30°C"

	// _, iconH := icon.GlyphSize()
	// top := r.Height() - iconH - margin

	// // Render text at left side
	// _, h := f.Measure(str)
	// f.Render(r, margin, top+(iconH-h), str)

	// // Render icon at right size
	// g := icon.GetGlyph('5')
	// glyphW, _ := g.GlyphSize()
	// g.Render(r, r.Width()-glyphW-margin, top)
}
