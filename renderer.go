package main

import "github.com/nsf/termbox-go"

type Renderer interface {
	Width() int
	Height() int

	Clear(fg, bg termbox.Attribute)

	SetFg(x, y int, attr termbox.Attribute)
	SetBg(x, y int, attr termbox.Attribute)
	SetCh(x, y int, ch rune)

	Commit() error
}

const Bg = termbox.ColorBlack
const Fg = termbox.ColorBlue
const FgFilled = termbox.ColorCyan | termbox.AttrBold
const ChEmpty = 'X'

type renderer struct {
	w, h int
	cs   []termbox.Cell
}

func NewRenderer() Renderer {
	w, h := termbox.Size()

	cells := make([]termbox.Cell, w*h)
	for i := range cells {
		cells[i] = termbox.Cell{Bg: Bg, Fg: Fg, Ch: ChEmpty}
	}

	return &renderer{w, h, cells}
}

func (r *renderer) Width() int  { return r.w }
func (r *renderer) Height() int { return r.h }

func (r *renderer) Commit() error {
	for i := 0; i < r.w; i++ {
		x := i
		for j := 0; j < r.h; j++ {
			y := j

			idx := r.calcIndex(i, j)
			cell := r.cs[idx]
			termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
		}
	}

	return termbox.Flush()
}

func (r *renderer) Clear(fg, bg termbox.Attribute) {
	for i := range r.cs {
		r.cs[i].Fg = fg
		r.cs[i].Bg = bg
	}
}

func (r *renderer) calcIndex(x, y int) int {
	i := r.w*y + x
	return i
}

func (r *renderer) SetFg(x, y int, attr termbox.Attribute) {
	i := r.calcIndex(x, y)
	r.cs[i].Fg = attr
}

func (r *renderer) SetBg(x, y int, attr termbox.Attribute) {
	i := r.calcIndex(x, y)
	r.cs[i].Bg = attr
}

func (r *renderer) SetCh(x, y int, ch rune) {
	i := r.calcIndex(x, y)
	r.cs[i].Ch = ch
}
