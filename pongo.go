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
		pollEvents(screen)
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

	defaultStyle := DefaultStyle()
	screen.SetStyle(defaultStyle)

	return screen, nil
}

// Listens for user input events, like keyboard input.
func pollEvents(screen tcell.Screen) {
	switch event := screen.PollEvent().(type) {
	case *tcell.EventResize:
		screen.Sync()
	case *tcell.EventKey:
		if isExitKey(event.Key()) {
			screen.Fini()
			os.Exit(0)
		}
	}
}

func isExitKey(key tcell.Key) bool {
	return key == tcell.KeyEscape || key == tcell.KeyCtrlC
}
