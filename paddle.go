package main

import "strings"

type Paddle struct {
	X      int
	Y      int
	YSpeed int
	Width  int
	Height int
}

// Displays the paddle as a stack of white spaces.
func (p *Paddle) Display() string {
	return strings.Repeat(" ", p.Height)
}
