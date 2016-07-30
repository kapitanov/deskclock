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
var iconFont Font

func init() {
	parseColor := func(str string, def termbox.Attribute) termbox.Attribute {
		switch str {
		case "cyan":
			return termbox.ColorCyan
		case "blue":
			return termbox.ColorBlue
		case "green":
			return termbox.ColorGreen
		case "magenta":
			return termbox.ColorMagenta
		case "red":
			return termbox.ColorRed
		case "white":
			return termbox.ColorWhite
		case "yellow":
			return termbox.ColorYellow
		case "black":
			return termbox.ColorBlack
		}
		return def
	}

	var fill int
	flag.IntVar(&fill, "char", 0x25C7, "Fill character (ASCII code)")
	flag.StringVar(&city, "city", "", "City ID")

	var fg, bg string
	flag.StringVar(&fg, "fg", "", "Foreground")
	flag.StringVar(&bg, "bg", "", "Background")

	flag.Parse()
	ChEmpty = rune(fill)
	Fg = parseColor(bg, termbox.ColorBlue)
	FgFilled = parseColor(fg, termbox.ColorCyan) | termbox.AttrBold
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

	iconFont, err = LoadFont("icon.ini")
	if err != nil {
		panic(err)
	}

	err = InitializeClock()
	if err != nil {
		panic(err)
	}

	err = InitializeWeather()
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
	DrawDate(r, lcdFont)
	DrawWeather(r, lcdFont, iconFont)

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
