package main

import "github.com/gdamore/tcell"

// The default style for the game screen.
func DefaultStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
}
