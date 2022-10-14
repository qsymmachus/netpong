package main

import (
	"time"

	"github.com/gdamore/tcell"
)

// Models a game of pong.
type Game struct {
	Screen  tcell.Screen
	Ball    Ball
	Player1 Paddle
	Player2 Paddle
}

// Starts the game.
func (g *Game) Run() {
	style := DefaultGameStyle()
	paddleStyle := DefaultPaddleStyle()

	for {
		g.Screen.Clear()
		width, height := g.Screen.Size()

		DrawSprite(
			g.Screen,
			g.Player1.X,
			g.Player1.Y,
			g.Player1.X+g.Player1.Width,
			g.Player1.Y+g.Player1.Height,
			paddleStyle,
			g.Player1.Display(),
		)

		DrawSprite(
			g.Screen,
			g.Player2.X,
			g.Player2.Y,
			g.Player2.X+g.Player1.Width,
			g.Player2.Y+g.Player1.Height,
			paddleStyle,
			g.Player2.Display(),
		)

		g.Ball.CheckEdges(width, height)
		g.Ball.Update()
		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, style)

		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}

// Draws a sprite on the screen, a group of runes with rectangular boundaries set
// by `xStart`, `yStart`, `xEnd`, and `yStart`.
func DrawSprite(s tcell.Screen, xStart, yStart, xEnd, yEnd int, style tcell.Style, text string) {
	row := yStart
	col := xStart

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++

		// If we've reached the vertical edge (last column), move down a row and
		// start from the first column.
		if col >= xEnd {
			row++
			col = xStart
		}

		// If we've reach the horizontal edge (last row), we've finished drawing
		// the sprite.
		if row > yEnd {
			break
		}
	}
}
