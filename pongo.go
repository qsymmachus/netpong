package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := initScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %+v\n", err)
	}

	width, height := screen.Size()

	ball := Ball{
		X:      width / 2,
		Y:      1,
		Xspeed: 1 * randomizeBallDirection(),
		Yspeed: 1,
	}

	player1 := Player{
		Score: 0,
		Paddle: Paddle{
			Width:  1,
			Height: 6,
			X:      5,
			Y:      (height / 2) - 3,
			YSpeed: 3,
		},
	}

	player2 := Player{
		Score: 0,
		Paddle: Paddle{
			Width:  1,
			Height: 6,
			X:      width - 5,
			Y:      (height / 2) - 3,
			YSpeed: 3,
		},
	}

	game := Game{
		Screen:      screen,
		Ball:        ball,
		LeftPlayer:  player1,
		RightPlayer: player2,
		MaxScore:    5,
	}

	go game.Run()

	for {
		pollEvents(&game)
	}
}

// Sets up the game screen.
func initScreen() (screen tcell.Screen, err error) {
	screen, err = tcell.NewScreen()
	if err != nil {
		return screen, err
	}

	if err := screen.Init(); err != nil {
		return screen, err
	}

	defaultStyle := DefaultGameStyle()
	screen.SetStyle(defaultStyle)

	return screen, nil
}

// Listens for user input events, like keyboard input.
func pollEvents(game *Game) {
	screen := game.Screen
	_, height := screen.Size()

	switch event := screen.PollEvent().(type) {
	case *tcell.EventResize:
		screen.Sync()
	case *tcell.EventKey:
		if isExitKey(event.Key()) {
			screen.Fini()
			os.Exit(0)
		} else if event.Rune() == 'w' {
			game.LeftPlayer.Paddle.MoveUp()
		} else if event.Rune() == 's' {
			game.LeftPlayer.Paddle.MoveDown(height)
		} else if event.Key() == tcell.KeyUp {
			game.RightPlayer.Paddle.MoveUp()
		} else if event.Key() == tcell.KeyDown {
			game.RightPlayer.Paddle.MoveDown(height)
		}
	}
}

func isExitKey(key tcell.Key) bool {
	return key == tcell.KeyEscape || key == tcell.KeyCtrlC
}

// Randomizes the direction of the ball along the X-axis by returning -1 or 1.
func randomizeBallDirection() (direction int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(2)

	if n == 1 {
		direction = 1
	} else {
		direction = -1
	}

	return direction
}
