package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

var redrawRequest = make(chan int)
var glyphs Glyphs

func main() {
	var err error
	glyphs, err = LoadGlyphs("glyphs.ini")
	if err != nil {
		panic(err)
	}

	err = termbox.Init()
	if err != nil {
		log.Panicf("Unable to init termbox: %s\n", err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)

	loop()
}

func redraw() {
	go func() {
		redrawRequest <- 0
	}()
}

func loop() {
	events := make(chan termbox.Event, 1)

	go func() {
		for {
			e := termbox.PollEvent()
			events <- e

			if e.Type == termbox.EventKey && e.Key == termbox.KeyEsc {
				return
			}
		}
	}()

	redraw()

	go timer()

	for {
		select {
		case <-redrawRequest:
			draw()
		case e := <-events:
			if e.Type == termbox.EventKey {
				if e.Key == termbox.KeyEsc {
					return
				}

				if e.Key == termbox.KeySpace {
					redraw()
				}
			}

			if e.Type == termbox.EventResize {
				redraw()
			}
		}
	}
}

func draw() {
	termbox.Sync()

	r := NewRenderer()

	h, m, _ := time.Now().Local().Clock()
	chars := make([]rune, 5)

	str := fmt.Sprintf("%02d", h)
	chars[0] = rune(str[0])
	chars[1] = rune(str[1])
	chars[2] = ':'
	str = fmt.Sprintf("%02d", m)
	chars[3] = rune(str[0])
	chars[4] = rune(str[1])

	xOffset := (r.Width() - 8*len(chars)) / 2
	yOffset := (r.Height() - 8) / 2

	if xOffset < 0 {
		xOffset = 0
	}
	if yOffset < 0 {
		yOffset = 0
	}

	for i := range chars {
		glyphs[chars[i]].Render(r, 8*i+xOffset, yOffset)
	}

	r.Commit()
}

func timer() {
	t := time.NewTicker(time.Minute)
	for {
		select {
		case <-t.C:
			redraw()
		}
	}
}
