package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var city string
var weather *Weather
var weatherLock sync.Mutex

func InitializeWeather() error {
	weather = GetWeather()

	go func() {
		t := time.NewTicker(10 * time.Minute)
		for {
			select {
			case <-t.C:
				w := GetWeather()

				weatherLock.Lock()
				weather = w
				weatherLock.Unlock()
			}
		}
	}()

	return nil
}

func DrawWeather(r Renderer, f Font, icon Font) {
	const (
		margin = 1
	)

	// icons:
	// 0 - rain
	// 1 - cloud
	// 2 - sun
	// 3 - storm
	// 4 - snow
	// 5 - heavy rain

	weatherLock.Lock()
	cw := weather
	weatherLock.Unlock()

	if cw == nil {
		return
	}

	sgn := "+"
	if cw.Temp < 0 {
		sgn = "-"
	} else if cw.Temp == 0 {
		sgn = " "
	}
	str := fmt.Sprintf("%s%02d°C", sgn, cw.Temp)

	w, h := f.Measure(str)
	f.Render(r, (r.Width()-w)/2, r.Height()-h-margin, str)

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

type Weather struct {
	Temp int
}

func GetWeather() *Weather {
	if city == "" {
		return nil
	}

	url := "https://pogoda.yandex.ru/" + city

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil
	}

	weatherElem := doc.Find(".current-weather__col_type_now").First()

	tempElem := weatherElem.Find(".current-weather__thermometer").First()
	tempTxt := tempElem.Text()
	matches := regexp.MustCompile("([\\+\\-][0-9]+)").FindStringSubmatch(tempTxt)
	if len(matches) <= 0 {
		return nil
	}

	temp, err := strconv.Atoi(matches[0])
	if err != nil {
		return nil
	}

	return &Weather{temp}
}
