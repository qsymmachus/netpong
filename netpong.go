package main

import (
	"flag"
	"log"

	"github.com/gdamore/tcell"
)

var (
	serverMode    = flag.Bool("server", false, "Host a netpong game as a server")
	port          = flag.Int("port", 60049, "The server port")
	serverAddress = flag.String("address", "localhost:60049", "The address of the netpong game to connect to in the format of host:port")
	DebugMode     = flag.Bool("debug", false, "When enabled, the game will log network errors encountered during play when it exits")
)

func main() {
	flag.Parse()

	screen, err := initScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %+v\n", err)
	}

	game := createGame(screen, *serverMode, *port, *serverAddress)
	game.Run()
}

func createGame(screen tcell.Screen, serverMode bool, port int, serverAddress string) Game {
	width, height := screen.Size()

	ball := Ball{
		X:      width / 2,
		Y:      1,
		Xspeed: 1 * ballDirection(serverMode),
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

	return Game{
		Screen:       screen,
		Ball:         ball,
		LocalPlayer:  player1,
		RemotePlayer: player2,
		MaxScore:     5,

		ServerMode:    serverMode,
		Port:          port,
		ServerAddress: serverAddress,

		Errors:    make(chan error, 0),
		DebugMode: *DebugMode,
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

func isExitKey(key tcell.Key) bool {
	return key == tcell.KeyEscape || key == tcell.KeyCtrlC
}

// Determines the initial direction of the ball along X-axis by returning -1 or 1.
// When the current player is running in server mode, the balls moves toward them,
// otherwise it moves away.
func ballDirection(serverMode bool) int {
	if serverMode {
		return -1
	} else {
		return 1
	}
}
