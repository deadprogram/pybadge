package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"github.com/conejoninja/pybadge/fonts"
	"tinygo.org/x/tinyfont"
)

const GRAVITY = 0.16

type FlappyGame struct {
	score       uint32
	status      uint8
	tubesHeight [4]int16
}

func (game *FlappyGame) Start() {
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
				if j > 15 {
					gopherBlack[j][i] = colors[RED]
				} else {
					gopherBlack[j][i] = colors[BLACK]
				}
			} else {
				gopherBlack[j][i] = colors[SKY]
			}
		}
	}

	play := true
	scoreStr := []byte("SCORE: 12345")
	for play {
		switch game.status {
		case START:
			display.SetScroll(0)
			display.StopScroll()
			display.FillScreen(colors[SKY])

			tinyfont.WriteLine(&display, &fonts.Bold12pt7b, 30, 50, []byte("FLAPPY"), colors[BLACK])
			tinyfont.WriteLine(&display, &fonts.Regular12pt7b, 8, 100, []byte("Press START"), colors[BLACK])
			for i := int16(0); i < 22; i++ {
				display.FillRectangleWithBuffer(124+i, 48, 1, 21, gopherBlack[i])
			}

			game.score = 0
			time.Sleep(2 * time.Second)
			for game.status == START {
				pressed, _ := buttons.Read8Input()
				if pressed&machine.BUTTON_START_MASK > 0 {
					game.status = PLAY
				}
				if pressed&machine.BUTTON_SELECT_MASK > 0 {
					game.status = QUIT
				}

			}
			break
		case GAMEOVER:
			display.SetScroll(0)
			display.StopScroll()
			display.FillScreen(colors[BLACK])

			scoreStr[7] = 48 + uint8(game.score/10000)
			scoreStr[8] = 48 + uint8((game.score/1000)%10)
			scoreStr[9] = 48 + uint8((game.score/100)%10)
			scoreStr[10] = 48 + uint8((game.score/10)%10)
			scoreStr[11] = 48 + uint8(game.score%10)

			tinyfont.WriteLine(&display, &fonts.Regular12pt7b, 8, 50, []byte("GAME OVER"), colors[TEXT])
			tinyfont.WriteLine(&display, &fonts.Regular12pt7b, 8, 100, []byte("Press START"), colors[TEXT])
			tinyfont.WriteLine(&display, &tinyfont.TomThumb, 50, 120, scoreStr, colors[TEXT])

			time.Sleep(2 * time.Second)
			for game.status == GAMEOVER {
				pressed, _ := buttons.Read8Input()
				if pressed&machine.BUTTON_START_MASK > 0 {
					game.status = START
				}
				if pressed&machine.BUTTON_SELECT_MASK > 0 {
					game.status = QUIT
				}

			}
			break
		case PLAY:
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
					game.status = QUIT
					break
				}

				if pressed&machine.BUTTON_A_MASK > 0 || pressed&machine.BUTTON_B_MASK > 0 {
					if !gPressed {
						speedY -= 4
						gPressed = true
					}
				} else {
					gPressed = false
				}

				game.flappyResetLine(159 - i)
				if (i > 10 && i < 20) || (i > 50 && i < 60) || (i > 90 && i < 100) || (i > 130 && i < 140) {
					if i == 19 {
						game.tubesHeight[0] = int16(rand.Int31n(30))
					}
					if i == 59 {
						game.tubesHeight[1] = int16(rand.Int31n(30))
					}
					if i == 99 {
						game.tubesHeight[2] = int16(rand.Int31n(30))
					}
					if i == 139 {
						game.tubesHeight[3] = int16(rand.Int31n(30))
					}
					game.flappyDrawTube(i)
				}
				if i == 15 || i == 55 || i == 95 || i == 135 {
					game.score++
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
				if gy > 87 {
					gy = 87
				}

				gx -= 1
				if gx < 0 {
					gx += 159
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
				for lastGy = 0; lastGy < 16; lastGy++ {
					display.FillRectangleWithBuffer(gx, gy, 1, 21, gopherBlack[lastGy])
					gx++
					if gx > 159 {
						gx -= 159
					}
				}
				lastGy = gy
				if game.checkCollisions(i+16, gy) && game.score > 4 {
					game.score -= 4
					game.status = GAMEOVER
					time.Sleep(120 * time.Second)
					break
				}

				display.SetScroll(i)
				i--
				if i < 0 {
					i = 159
				}
				time.Sleep(4 * time.Millisecond)
			}
			break
		case QUIT:
			display.FillScreen(colors[BLACK])
			play = false
			break
		}
	}
	display.SetScroll(0)
	display.StopScroll()
}

func (g *FlappyGame) flappyResetLine(i int16) {
	display.FillRectangle(i, 0, 1, 108, colors[SKY])
	display.FillRectangle(i, 108, 1, 2, colors[DARKBROWN])
	display.FillRectangle(i, 110, 1, 18, colors[BROWN])
	//display.FillRectangle(i+1, 0, 1, 128, colors[RED])
}

func (g *FlappyGame) flappyDrawTube(i int16) {
	h := g.tubesHeight[3]
	if i > 10 && i < 20 {
		h = g.tubesHeight[0]
	} else if i > 50 && i < 60 {
		h = g.tubesHeight[0]
	} else if i > 90 && i < 100 {
		h = g.tubesHeight[0]
	}
	i = 159 - i
	display.FillRectangle(i, 0, 1, 10+h, colors[GREEN])
	display.FillRectangle(i, 70+h, 1, 58-h, colors[GREEN])
	if i == 159-19 {
		display.FillRectangle(i, 0, 1, 10+h, colors[RED])
	}
	if i == 159-59 {
		display.FillRectangle(i, 0, 1, 10+h, colors[BLACK])
	}
	if i == 159-99 {
		display.FillRectangle(i, 0, 1, 10+h, colors[DARKBROWN])
	}
	if i == 159-139 {
		display.FillRectangle(i, 0, 1, 10+h, colors[GREEN])
	}
}

func (g *FlappyGame) checkCollisions(i, gy int16) bool {
	if i > 159 {
		i = i-159
	}
	if ((i > 11 && i < 19) && (gy < g.tubesHeight[3]+9 || gy > g.tubesHeight[3]+50)) ||
		((i > 51 && i < 59) && (gy < g.tubesHeight[0]+9 || gy > g.tubesHeight[0]+50)) ||
		((i > 91 && i < 99) && (gy < g.tubesHeight[1]+9 || gy > g.tubesHeight[1]+50)) ||
		((i > 131 && i < 139) && (gy < g.tubesHeight[2]+9 || gy > g.tubesHeight[2]+50)) {
		println(i, gy, g.tubesHeight[0]+9, g.tubesHeight[0]+50)
		println(i, gy, g.tubesHeight[1]+9, g.tubesHeight[1]+50)
		println(i, gy, g.tubesHeight[2]+9, g.tubesHeight[2]+50)
		println(i, gy, g.tubesHeight[3]+9, g.tubesHeight[3]+50)
		return true
	}
	return false
}
