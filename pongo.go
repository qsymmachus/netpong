package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, style, err := initScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %+v\n", err)
	}

	for {
		drawContent(screen, style)
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

// Draws the game content.
func drawContent(screen tcell.Screen, style tcell.Style) {
	screen.SetContent(0, 0, 'H', nil, style)
	screen.SetContent(1, 0, 'i', nil, style)
	screen.SetContent(2, 0, '!', nil, style)

	screen.Show()
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
