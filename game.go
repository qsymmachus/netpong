package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

// Models a game of pong.
type Game struct {
	Screen      tcell.Screen
	Ball        Ball
	LeftPlayer  Player
	RightPlayer Player
	MaxScore    int
}

// Starts the game.
func (g *Game) Run() {
	style := DefaultGameStyle()
	paddleStyle := DefaultPaddleStyle()

	// Continually poll for events (game inputs and outputs) in the background.
	// Events update game state.
	go func() {
		for {
			g.PollEvents()
		}
	}()

	// Control loop that continually checks game state and redraws the screen based
	// on that state.
	for {
		g.Screen.Clear()

		if g.GameOver() {
			g.DrawEndGame(style)
			break
		}

		g.DrawPaddles(paddleStyle)
		g.DrawBall(style)
		g.DrawScores(style)

		pause(40)
		g.Screen.Show()
	}
}

// Draw the player paddles on the game screen.
func (g *Game) DrawPaddles(paddleStyle tcell.Style) {
	drawSprite(
		g.Screen,
		g.LeftPlayer.Paddle.X,
		g.LeftPlayer.Paddle.Y,
		g.LeftPlayer.Paddle.X+g.LeftPlayer.Paddle.Width,
		g.LeftPlayer.Paddle.Y+g.LeftPlayer.Paddle.Height,
		paddleStyle,
		g.LeftPlayer.Paddle.Display(),
	)

	drawSprite(
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

	drawSprite(g.Screen, g.Ball.X, g.Ball.Y, g.Ball.X, g.Ball.Y, style, g.Ball.Display())
}

// Draws the player scores on the game screen.
func (g *Game) DrawScores(style tcell.Style) {
	width, _ := g.Screen.Size()
	leftScore := fmt.Sprintf("← %s", strconv.Itoa(g.LeftPlayer.Score))
	rightScore := fmt.Sprintf("%s →", strconv.Itoa(g.RightPlayer.Score))

	drawSprite(g.Screen, (width/2)-10, 1, (width/2)-7, 1, style, leftScore)
	drawSprite(g.Screen, (width/2)+10, 1, (width/2)+13, 1, style, rightScore)
}

// Checks if the game is over.
func (g *Game) GameOver() bool {
	return g.LeftPlayer.Score == g.MaxScore || g.RightPlayer.Score == g.MaxScore
}

// Declares the winner of the game.
func (g *Game) DeclareWinner() string {
	if g.LeftPlayer.Score > g.RightPlayer.Score {
		return "← Winner"
	} else {
		return "Winner →"
	}
}

// Listens for user input events, like keyboard input.
func (g *Game) PollEvents() {
	screen := g.Screen
	_, height := screen.Size()

	switch event := screen.PollEvent().(type) {
	case *tcell.EventResize:
		screen.Sync()
	case *tcell.EventKey:
		if isExitKey(event.Key()) {
			screen.Fini()
			os.Exit(0)
		} else if event.Rune() == 'w' {
			g.LeftPlayer.Paddle.MoveUp()
		} else if event.Rune() == 's' {
			g.LeftPlayer.Paddle.MoveDown(height)
		} else if event.Key() == tcell.KeyUp {
			g.RightPlayer.Paddle.MoveUp()
		} else if event.Key() == tcell.KeyDown {
			g.RightPlayer.Paddle.MoveDown(height)
		}
	}
}

func (g *Game) DrawEndGame(style tcell.Style) {
	width, _ := g.Screen.Size()

	drawSprite(g.Screen, (width/2)-4, 7, (width/2)+5, 7, style, g.DeclareWinner())
	drawSprite(g.Screen, (width/2)-8, 10, (width/2)+8, 10, style, "(CTRL+C to exit)")
	g.Screen.Show()
}

// Draws a sprite on the screen, a group of runes with rectangular boundaries set
// by `xStart`, `yStart`, `xEnd`, and `yStart`.
func drawSprite(s tcell.Screen, xStart, yStart, xEnd, yEnd int, style tcell.Style, text string) {
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
