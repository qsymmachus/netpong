package main

import "github.com/gdamore/tcell"

// The default style for the game screen.
func DefaultGameStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
}

// The default style for a paddle.
func DefaultPaddleStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
}
