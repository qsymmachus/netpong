package main

import (
	"strings"
)

type Paddle struct {
	X      int
	Y      int
	YSpeed int
	Width  int
	Height int
}

// Moves the position of the paddle up.
func (p *Paddle) MoveUp() {
	if p.Y > 0 {
		p.Y -= p.YSpeed
	}
}

// Moves the position of the paddle down.
func (p *Paddle) MoveDown(windowHeight int) {
	if p.Y < windowHeight-p.Height {
		p.Y += p.YSpeed
	}
}

// Displays a paddle as a stack of white spaces.
func (p *Paddle) Display() string {
	return strings.Repeat(" ", p.Height)
}
