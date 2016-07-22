package main

import (
	"flag"
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

var redrawRequest = make(chan int)
var mainFont Font
var lcdFont Font

func init() {
	var fill int
	flag.IntVar(&fill, "char", 0x25C7, "Fill character (ASCII code)")

	flag.Parse()

	ChEmpty = rune(fill)
}

func main() {
	var err error

	mainFont, err = LoadFont("font.ini")
	if err != nil {
		panic(err)
	}

	lcdFont, err = LoadFont("lcd.ini")
	if err != nil {
		panic(err)
	}

	err = InitializeClock()
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

	DrawClock(r, mainFont)
	DrawDate(r, mainFont)
	/*
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
			font[chars[i]].Render(r, 8*i+xOffset, yOffset)
		}

		// Draw title

		title := "CLOCK"
		yOffset = 2
		xOffset = (r.Width() - 5*len(title)) / 2

		for i, c := range title {
			lcd[c].Render(r, 5*i+xOffset, yOffset)
		}

		// Draw date

		_, month, day := time.Now().Local().Date()
		date := fmt.Sprintf("%02d.%02d", month, day)
		yOffset = r.Height() - 5 - 2
		xOffset = (r.Width() - 5*len(title)) / 2

		for i, c := range date {
			lcd[c].Render(r, 5*i+xOffset, yOffset)
		}*/

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
