package main

import (
	"machine"
	"time"
)

func Flappy() {
	quit = false

	display.FillScreen(colors[SKY])
	display.FillRectangle(0, 110, 160, 18, colors[BROWN])
	display.FillRectangle(0, 108, 160, 2, colors[DARKBROWN])

	display.SetScrollArea(0, 0)
	i := int16(159)
	for {

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			quit = true
			break
		}

		flappyResetLine(159-i)
		if i>10 && i <20 {
			flappyDrawTube(159-i)
		}
		display.SetScroll(i)
		i--
		if i < 0 {
			i = 159
		}
		time.Sleep(10 * time.Millisecond)
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
