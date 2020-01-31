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
	var gopherBlack [22][]color.RGBA
	for j := 0; j < 22; j++ {
		gopherBlack[j] = make([]color.RGBA, 21)
		for i := 0; i < 21; i++ {
			if gopherBalloon[i][j] {
				gopherBlack[j][i] = colors[BLACK]
			} else {
				gopherBlack[j][i] = colors[SKY]
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
	var k int16
	speedY := float32(-1)
	gy = 60
	lastGy = 60
	gPressed := false
	for {

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			time.Sleep(120 * time.Second)
			quit = true
			break
		}

		if pressed&machine.BUTTON_A_MASK > 0 || pressed&machine.BUTTON_B_MASK > 0 {
			if !gPressed {
				speedY -= 4
				gPressed = true
			} else {
				speedY -= GRAVITY * 0.6
			}
		} else {
			gPressed = false
		}

		flappyResetLine(159 - i)
		if (i > 10 && i < 20) || (i > 140 && i < 150) {
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
		gy = gy + int16(speedY)
		if speedY < -3 {
			speedY = -3
		}

		if gy < 0 {
			gy = 0
		}
		if gy > 86 {
			gy = 86
		}

		gx -=1
		if gx <0 {
			gx+=159
		}
		for k = 0; k < 22; k++ {
			display.FillRectangle(gx, lastGy, 1, 21, colors[SKY])
			gx++
			if gx > 159 {
				gx -= 159
			}
		}
		gx -= 21
		if gx < 0 {
			gx += 159
		}
		for lastGy = 0; lastGy < 22; lastGy++ {
			display.FillRectangleWithBuffer(gx, gy, 1, 21, gopherBlack[lastGy])
			gx++
			if gx > 159 {
				gx -= 159
			}
		}
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
		time.Sleep(4 * time.Millisecond)
	}
	display.SetScroll(0)
	display.StopScroll()
}

func flappyResetLine(i int16) {
	display.FillRectangle(i, 0, 1, 108, colors[SKY])
	display.FillRectangle(i, 108, 1, 2, colors[DARKBROWN])
	display.FillRectangle(i, 110, 1, 18, colors[BROWN])
	display.FillRectangle(i+1, 0, 1, 128, colors[RED])
}

func flappyDrawTube(i int16) {
	display.FillRectangle(i, 80, 1, 128-80, colors[GREEN])
}
