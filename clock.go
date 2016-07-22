package main

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

func DrawClock(r Renderer, f Font) {
	h, m, _ := time.Now().Local().Clock()

	str := fmt.Sprintf("%02d:%02d", h, m)

	w, h := f.Measure(str)
	f.Render(r, (r.Width()-w)/2, (r.Height()-h)/2, str)
}

func DrawDate(r Renderer, f Font) {
	_, month, day := time.Now().Local().Date()
	dayOfWeek := GetDayOfWeek(time.Now().Local().Weekday())
	str := fmt.Sprintf("%s %2d.%02d", dayOfWeek, month, day)

	w, _ := f.Measure(str)
	f.Render(r, (r.Width()-w)/2, 1, str)
}

var weekdayNames = make(map[time.Weekday]string)

func InitializeClock() error {
	bytes, err := Asset("res/res.ini")
	if err != nil {
		return err
	}

	file, err := ini.Load(bytes)
	if err != nil {
		return err
	}

	section := file.Section("")
	weekdayNames[time.Monday] = section.Key("Monday").Value()
	weekdayNames[time.Tuesday] = section.Key("Tuesday").Value()
	weekdayNames[time.Wednesday] = section.Key("Wednesday").Value()
	weekdayNames[time.Thursday] = section.Key("Thursday").Value()
	weekdayNames[time.Friday] = section.Key("Friday").Value()
	weekdayNames[time.Saturday] = section.Key("Saturday").Value()
	weekdayNames[time.Sunday] = section.Key("Sunday").Value()

	return nil
}

func GetDayOfWeek(w time.Weekday) string {
	str, exists := weekdayNames[w]
	if !exists {
		return "?"
	}

	return str
}
