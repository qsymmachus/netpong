package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// pb "github.com/qsymmachus/netpong/netpong"
)

var (
	serverAddress = flag.String("address", "localhost:60049", "The address of the netpong game to connect to in the format of host:port")
)

func main() {
	flag.Parse()

	connOpts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(*serverAddress, connOpts)
	if err != nil {
		log.Fatalf("Failed to connect to remote game: %v\n", err)
	}
	defer conn.Close()

	// TODO do something with the client
	// client := pb.NewNetPongClient(conn)

	screen, err := initScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %+v\n", err)
	}

	game := createGame(screen)
	game.Run()
}

func createGame(screen tcell.Screen) Game {
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

	return Game{
		Screen:      screen,
		Ball:        ball,
		LeftPlayer:  player1,
		RightPlayer: player2,
		MaxScore:    5,
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
