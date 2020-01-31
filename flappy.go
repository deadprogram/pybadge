package main

import (
	"image/color"
	"machine"
	"time"
)

const GRAVITY = 0.1

func Flappy() {
	gopherBalloon := [21][22]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, true, false},
		{false, false, true, true, true, false, false, false, false, false, false, true, true, true, false, false, true, false, false, false, false, true},
		{false, false, true, false, false, true, true, true, true, true, true, false, false, true, false, true, false, false, false, false, false, true},
		{false, false, true, true, false, false, false, false, false, false, false, false, false, true, false, false, true, false, false, false, true, false},
		{false, true, false, false, true, false, false, false, false, true, true, true, true, true, false, false, true, false, false, false, true, false},
		{true, true, true, false, false, true, false, false, true, false, false, true, true, false, true, false, false, true, false, true, false, false},
		{true, true, true, false, false, true, false, false, true, false, false, true, true, false, true, false, false, false, true, false, false, false},
		{true, false, false, false, false, true, false, false, true, false, false, false, false, false, true, false, false, true, false, false, false, false},
		{true, false, false, false, true, false, false, false, false, true, false, false, false, true, true, false, false, true, false, false, false, false},
		{false, true, false, false, true, false, true, true, false, false, true, true, true, false, true, false, true, false, false, false, false, false},
		{false, true, true, true, false, false, false, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false},
		{true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false},
		{true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false},
		{false, false, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	}
	gopherBlack := make([]color.RGBA, 462)
	for i := 0; i < 21; i++ {
		for j := 0; j < 22; j++ {
			if gopherBalloon[i][j] {
				gopherBlack[22*i+j] = colors[BLACK]
			} else {
				gopherBlack[22*i+j] = colors[SKY]
			}
		}
	}

	quit = false

	display.FillScreen(colors[SKY])
	display.FillRectangle(0, 110, 160, 18, colors[BROWN])
	display.FillRectangle(0, 108, 160, 2, colors[DARKBROWN])

	display.SetScrollArea(0, 0)
	i := int16(159)
	var gx int16
	var gy int16
	var lastGy int16
	speedY := float32(-1)
	gy = 60
	lastGy = 60
	for {

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			time.Sleep(120 * time.Second)
			quit = true
			break
		}

		if pressed&machine.BUTTON_A_MASK > 0 || pressed&machine.BUTTON_B_MASK > 0 {
			speedY -= 2
		}

		flappyResetLine(159 - i + 1)
		if i > 10 && i < 20 {
			flappyDrawTube(159 - i)
		}

		// draw Gopher
		gx = 169 - i
		if gx > 159 {
			gx -= 159
		}
		speedY += GRAVITY
		if speedY > 1.4 {
			speedY = 1.4
		}
		if speedY < -1.4 {
			speedY = -1.4
		}

		gy = gy + int16(speedY)
		display.FillRectangle(gx-1, lastGy, 22, 21, colors[SKY])
		display.FillRectangleWithBuffer(gx, gy, 22, 21, gopherBlack)
		//tinyfont.DrawChar(&display, &fonts.Regular32pt, gx-1, lastGy, byte('E'), colors[SKY])
		//tinyfont.DrawChar(&display, &fonts.Regular32pt, gx, gy, byte('E'), colors[BLACK])
		//display.FillRectangle(gx-1, lastGy, 14, 14, colors[SKY])
		//display.FillRectangle(gx-1, gy, 14, 14, colors[BLACK])
		lastGy = int16(gy)

		display.SetScroll(i)
		i--
		if i < 0 {
			i = 159
		}
		time.Sleep(1 * time.Millisecond)
	}
	display.SetScroll(0)
	display.StopScroll()
}

func flappyResetLine(i int16) {
	display.FillRectangle(i, 0, 1, 108, colors[SKY])
	display.FillRectangle(i, 108, 1, 2, colors[DARKBROWN])
	display.FillRectangle(i, 110, 1, 18, colors[BROWN])
}

func flappyDrawTube(i int16) {
	display.FillRectangle(i, 80, 1, 128-80, colors[GREEN])
}
