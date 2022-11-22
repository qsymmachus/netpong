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

// Checks if the ball has hit a top or bottom edge of the game screen.
// If it has hit the  edges, the ball "bounces" by reversing its direction
// of travel.
func (b *Ball) CheckEdges(maxWidth, maxHeight int) {
	if b.Y <= 0 || b.Y >= maxHeight {
		b.ReverseY()
	}
}

// Checks if the ball has collided with one of the players paddles. If it has,
// the ball "bounces" by reversing its current direction of travel.
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

// Resets the ball position with the ball traveling left.
func (b *Ball) ResetLeft(width int) {
	b.X = width / 2
	b.Y = 1
	b.Xspeed = -1
	b.Yspeed = 1
}

// Resets the ball position with the ball traveling right.
func (b *Ball) ResetRight(width int) {
	b.X = width / 2
	b.Y = 1
	b.Xspeed = 1
	b.Yspeed = 1
}

// Checks if the ball hit the left edge.
func (b *Ball) HasHitLeft() bool {
	return b.X <= 0
}

// Checks if the ball hit the right edge.
func (b *Ball) HasHitRight(width int) bool {
	return b.X >= width
}
