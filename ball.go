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
		b.ReverseX()
	}

	if b.Y <= 0 || b.Y >= maxHeight {
		b.ReverseY()
	}
}

// Checks if the ball has collided with one of the players paddles. If it has,
// the ball "bounces" by reverse its current direction of travel.
func (b *Ball) CheckCollisions(player1, player2 Paddle) {
	if b.Intersects(player1) || b.Intersects(player2) {
		b.ReverseX()
		b.ReverseY()
	}
}

// Checks if the ball has collided with a paddle.
func (b *Ball) Intersects(p Paddle) bool {
	return b.X >= p.X && b.X <= p.X+p.Width && b.Y >= p.Y && p.Y <= p.Y+p.Height
}

// Reverses the ball's direction along the X axis.
func (b *Ball) ReverseX() {
	b.Xspeed *= -1
}

// Reserves the ball's direction along the Y axis.
func (b *Ball) ReverseY() {
	b.Yspeed *= -1
}
