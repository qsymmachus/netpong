package main

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	screen, style, err := initScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %+v\n", err)
	}

	go run(screen, style)

	for {
		pollEvents(screen)
	}
}

// Sets up the game screen.
func initScreen() (screen tcell.Screen, style tcell.Style, err error) {
	screen, err = tcell.NewScreen()
	if err != nil {
		return screen, style, err
	}

	if err := screen.Init(); err != nil {
		return screen, style, err
	}

	defaultStyle := defaultStyle()
	screen.SetStyle(defaultStyle)

	return screen, defaultStyle, nil
}

// The default style for the screen.
func defaultStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
}

// Runs the game.
func run(screen tcell.Screen, style tcell.Style) {
	x := 0
	for {
		screen.Clear()

		screen.SetContent(x, 10, 'H', nil, style)
		screen.SetContent(x+1, 10, 'i', nil, style)
		screen.SetContent(x+2, 10, '!', nil, style)

		screen.Show()
		x++

		time.Sleep(40 * time.Millisecond)
	}
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
