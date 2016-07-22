//go:generate go-bindata -o res.go res/

package main

import (
	"errors"
	"fmt"

	"gopkg.in/ini.v1"
)

var ErrLetterSizeMismatch = errors.New("E_LETTER_SIZE_MISMATCH")

type Glyph interface {
	GlyphSize() (int, int)
	Render(r Renderer, x, y int)
	String() string
}

type Font interface {
	GlyphSize() (int, int)
	Runes() []rune
	GetGlyph(c rune) Glyph
	Measure(str string) (int, int)
	Render(r Renderer, x, y int, str string)
}

func LoadFont(path string) (Font, error) {
	glyphs := make(map[rune]*glyphImpl)

	var width, height int

	bytes, err := Asset("res/" + path)
	if err != nil {
		return nil, err
	}

	file, err := ini.Load(bytes)
	if err != nil {
		return nil, err
	}

	for _, section := range file.Sections() {
		glyph, err := parseGlyph(section, &width, &height)
		if err != nil {
			return nil, err
		}

		glyphs[glyph.c] = glyph
	}

	emptyGlyph := &glyphImpl{rune(0), make([][]bool, 0), width, height}

	return &fontImpl{width, height, glyphs, emptyGlyph}, nil
}

func parseGlyph(section *ini.Section, width, height *int) (*glyphImpl, error) {
	var lines [][]bool

	w := 0
	for _, key := range section.Keys() {
		line := parseGlyphLine(key.Value())
		if len(line) > w {
			w = len(line)
		}
		lines = append(lines, parseGlyphLine(key.Value()))
	}

	h := len(lines)

	if *height == 0 {
		*height = h
	} else if *height != h {
		return nil, ErrLetterSizeMismatch
	}

	var name rune
	for _, c := range section.Name() {
		name = c
		break
	}

	glyph := &glyphImpl{name, lines, w, h}
	return glyph, nil
}

func parseGlyphLine(s string) []bool {
	line := make([]bool, len(s))

	for i, c := range s {
		if c == '1' {
			line[i] = true
		}
	}

	return line
}

type glyphImpl struct {
	c     rune
	lines [][]bool
	w, h  int
}

func (g *glyphImpl) GlyphSize() (int, int) {
	return g.w, g.h
}

func (g *glyphImpl) String() string {
	s := ""
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			if g.lines[i][j] {
				s += "x"
			} else {
				s += "-"
			}
		}
		s += "\n"
	}

	return s
}

func (g *glyphImpl) Render(r Renderer, x, y int) {
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			if g.lines[i][j] {
				r.SetFg(x+j, y+i, FgFilled)
			} else {
				r.SetFg(x+j, y+i, Fg)
			}
		}
	}
}

type fontImpl struct {
	w, h       int
	glyphs     map[rune]*glyphImpl
	emptyGlyph *glyphImpl
}

func (f *fontImpl) GlyphSize() (int, int) {
	return f.w, f.h
}

func (f *fontImpl) Runes() []rune {
	runes := make([]rune, len(f.glyphs))
	i := 0
	for c := range f.glyphs {
		runes[i] = c
		i++
	}
	return runes
}

func (f *fontImpl) GetGlyph(c rune) Glyph {
	glyph, exists := f.glyphs[c]
	if !exists {
		panic(fmt.Sprintf("<%c>", c)) // todo
		return f.emptyGlyph
	}

	return glyph
}

func (f *fontImpl) Measure(str string) (int, int) {
	var width, height int

	for _, c := range str {
		glyph := f.GetGlyph(rune(c))
		w, h := glyph.GlyphSize()
		width += w

		if height < h {
			height = h
		}
	}

	return width, height
}

func (f *fontImpl) Render(r Renderer, x, y int, str string) {
	for _, c := range str {
		glyph := f.GetGlyph(rune(c))
		glyph.Render(r, x, y)

		w, _ := glyph.GlyphSize()
		x += w
	}
}
