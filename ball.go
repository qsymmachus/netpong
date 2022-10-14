package main

// The pong ball!
type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

// Displays the ball as a white dot.
func (b *Ball) Display() string {
	return "\u25CF"
}

// Updates the ball's position by incrementing its position by its speed.
func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

// Checks if the ball has hit an edge of the game screen. If it has, the ball
// "bounces" by reversing its current direction along the correct dimension.
func (b *Ball) CheckEdges(maxWidth, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
		b.Xspeed *= -1
	}

	if b.Y <= 0 || b.Y >= maxHeight {
		b.Yspeed *= -1
	}
}
