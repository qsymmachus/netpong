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
	style := DefaultStyle()

	for {
		g.Screen.Clear()

		width, height := g.Screen.Size()
		g.Ball.CheckEdges(width, height)

		g.Ball.Update()
		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, style)

		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
