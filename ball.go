package main

// The pong ball!
type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

// Displays the ball as a white dot.
func (b *Ball) Display() rune {
	return '\u25CF'
}

// Updates the ball's position by incrementing its position by its speed.
func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}
