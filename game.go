package main

import (
	"time"

	"github.com/gdamore/tcell"
)

// Models a game of pong.
type Game struct {
	Screen      tcell.Screen
	Ball        Ball
	LeftPlayer  Player
	RightPlayer Player
}

// Starts the game.
func (g *Game) Run() {
	style := DefaultGameStyle()
	paddleStyle := DefaultPaddleStyle()

	for {
		g.Screen.Clear()

		g.DrawPaddles(paddleStyle)
		g.DrawBall(style)

		pause(40)
		g.Screen.Show()
	}
}

// Draw the player paddles on the game screen.
func (g *Game) DrawPaddles(paddleStyle tcell.Style) {
	DrawSprite(
		g.Screen,
		g.LeftPlayer.Paddle.X,
		g.LeftPlayer.Paddle.Y,
		g.LeftPlayer.Paddle.X+g.LeftPlayer.Paddle.Width,
		g.LeftPlayer.Paddle.Y+g.LeftPlayer.Paddle.Height,
		paddleStyle,
		g.LeftPlayer.Paddle.Display(),
	)

	DrawSprite(
		g.Screen,
		g.RightPlayer.Paddle.X,
		g.RightPlayer.Paddle.Y,
		g.RightPlayer.Paddle.X+g.RightPlayer.Paddle.Width,
		g.RightPlayer.Paddle.Y+g.RightPlayer.Paddle.Height,
		paddleStyle,
		g.RightPlayer.Paddle.Display(),
	)
}

// Draw the ball on the game screen.
func (g *Game) DrawBall(style tcell.Style) {
	width, height := g.Screen.Size()

	g.Ball.CheckEdges(width, height)
	g.Ball.CheckCollisions(g.LeftPlayer.Paddle, g.RightPlayer.Paddle)

	if g.Ball.HasHitLeft() {
		g.RightPlayer.Score += 1
		pause(1000)
		g.Ball.ResetLeft(width)
	}

	if g.Ball.HasHitRight(width) {
		g.LeftPlayer.Score += 1
		pause(1000)
		g.Ball.ResetRight(width)
	}

	g.Ball.Update()

	DrawSprite(g.Screen, g.Ball.X, g.Ball.Y, g.Ball.X, g.Ball.Y, style, g.Ball.Display())
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

func pause(milliseconds time.Duration) {
	time.Sleep(milliseconds * time.Millisecond)
}
