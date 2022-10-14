package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := initScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %+v\n", err)
	}

	width, _ := screen.Size()

	ball := Ball{
		X:      1,
		Y:      1,
		Xspeed: 1,
		Yspeed: 1,
	}

	player1 := Paddle{
		Width:  1,
		Height: 6,
		X:      5,
		Y:      3,
		YSpeed: 3,
	}

	player2 := Paddle{
		Width:  1,
		Height: 6,
		X:      width - 5,
		Y:      3,
		YSpeed: 3,
	}

	game := Game{
		Screen:  screen,
		Ball:    ball,
		Player1: player1,
		Player2: player2,
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
			game.Player1.MoveUp()
		} else if event.Rune() == 's' {
			game.Player1.MoveDown(height)
		} else if event.Key() == tcell.KeyUp {
			game.Player2.MoveUp()
		} else if event.Key() == tcell.KeyDown {
			game.Player2.MoveDown(height)
		}
	}
}

func isExitKey(key tcell.Key) bool {
	return key == tcell.KeyEscape || key == tcell.KeyCtrlC
}
