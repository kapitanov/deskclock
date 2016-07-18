package main

import "gopkg.in/ini.v1"

type Glyph struct {
	lines [][]bool
	w, h  int
}

type Glyphs map[rune]*Glyph

func LoadGlyphs(path string) (Glyphs, error) {
	glyphs := make(Glyphs)

	file, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	for _, section := range file.Sections() {
		name := section.Name()
		var lines [][]bool

		w := 0
		for _, key := range section.Keys() {
			line := parseGlyphLine(key.Value())
			if len(line) > w {
				w = len(line)
			}
			lines = append(lines, parseGlyphLine(key.Value()))
		}

		c := rune(name[0])
		glyphs[c] = &Glyph{lines, w, len(lines)}
	}

	return glyphs, nil
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

func (g *Glyph) String() string {
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

func (g *Glyph) Render(r Renderer, x, y int) {
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			if g.lines[i][j] {
				r.SetFg(x+j, y+i, FgFilled)

				// r.SetCh(x+j, y+i, ChFill)
			} else {
				r.SetFg(x+j, y+i, Fg)
				// r.SetCh(x+j, y+i, ChEmpty)
			}
		}
	}
}
