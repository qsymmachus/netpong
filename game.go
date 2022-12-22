package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	pb "github.com/qsymmachus/netpong/netpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Models a game of pong. Implements the `NetPongServer` interface.
type Game struct {
	pb.UnimplementedNetPongServer

	Screen       tcell.Screen
	Ball         Ball
	LocalPlayer  Player
	RemotePlayer Player
	MaxScore     int

	Connected     bool
	ServerMode    bool
	Port          int
	ServerAddress string

	Errors    chan (error)
	DebugMode bool
}

// I created this interface so I could write a version of the `Play` function
// that accepts either a `NetPong_PlayServer` or `NetPong_PlayClient`.
type PlayStream interface {
	Send(*pb.MovePaddle) error
	Recv() (*pb.MovePaddle, error)
}

// Implementation of streamed gameplay between a local and remote player. The
// function is identical for clients and servers.
func (g *Game) PlayWithStream(stream PlayStream) error {
	g.Connected = true
	defer func() { g.Connected = false }()
	_, height := g.Screen.Size()

	for {
		// Receive paddle movement commands from the remote game and use
		// them to update local game state.
		remoteMove, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			g.Errors <- err
			return err
		}

		if remoteMove.Direction == pb.Direction_UP {
			g.RemotePlayer.Paddle.MoveUp()
		} else if remoteMove.Direction == pb.Direction_DOWN {
			g.RemotePlayer.Paddle.MoveDown(height)
		}

		// Poll local key events and turn them into paddle movement commands to
		// send them to the remote game.
		switch event := g.Screen.PollEvent().(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyUp {
				err := stream.Send(&pb.MovePaddle{Direction: pb.Direction_UP})
				if err != nil {
					g.Errors <- err
				}
			} else if event.Key() == tcell.KeyDown {
				err := stream.Send(&pb.MovePaddle{Direction: pb.Direction_DOWN})
				if err != nil {
					g.Errors <- err
				}
			}
		}
	}
}

// Server-side handling of a stream of paddle movement commands between
// the local and remote games.
func (g *Game) Play(stream pb.NetPong_PlayServer) error {
	return g.PlayWithStream(stream)
}

// Starts the game.
func (g *Game) Run() {
	style := DefaultGameStyle()
	paddleStyle := DefaultPaddleStyle()

	// Continually poll for events (game inputs and outputs) in the background.
	// Events update game state.
	go func() {
		for {
			g.PollEvents()
		}
	}()

	if g.ServerMode {
		// If the game is running in server mode, listen and wait for a connection.
		listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", g.Port))
		if err != nil {
			log.Fatalf("Failed to start game server: %v\n", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterNetPongServer(grpcServer, g)
		go grpcServer.Serve(listen)

		for !g.Connected {
			g.DrawWaitScreen(style)
		}
	} else {
		// Otherwise, the game is in client mode, and will attempt to connect to a server.
		connOpts := grpc.WithTransportCredentials(insecure.NewCredentials())
		conn, err := grpc.Dial(*serverAddress, connOpts)
		if err != nil {
			log.Fatalf("Failed to connect to remote game: %v\n", err)
		}
		defer conn.Close()

		client := pb.NewNetPongClient(conn)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		stream, err := client.Play(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to remote game: %v\n", err)
		}

		go g.PlayWithStream(stream)
	}

	// Control loop that continually checks game state and redraws the screen based
	// on that state.
	for {
		g.Screen.Clear()

		if g.GameOver() {
			g.DrawEndGame(style)
			break
		}

		g.DrawPaddles(paddleStyle)
		g.DrawBall(style)
		g.DrawScores(style)

		pause(40)
		g.Screen.Show()
	}
}

// Draw the player paddles on the game screen.
func (g *Game) DrawPaddles(paddleStyle tcell.Style) {
	drawSprite(
		g.Screen,
		g.LocalPlayer.Paddle.X,
		g.LocalPlayer.Paddle.Y,
		g.LocalPlayer.Paddle.X+g.LocalPlayer.Paddle.Width,
		g.LocalPlayer.Paddle.Y+g.LocalPlayer.Paddle.Height,
		paddleStyle,
		g.LocalPlayer.Paddle.Display(),
	)

	drawSprite(
		g.Screen,
		g.RemotePlayer.Paddle.X,
		g.RemotePlayer.Paddle.Y,
		g.RemotePlayer.Paddle.X+g.RemotePlayer.Paddle.Width,
		g.RemotePlayer.Paddle.Y+g.RemotePlayer.Paddle.Height,
		paddleStyle,
		g.RemotePlayer.Paddle.Display(),
	)
}

// Draw the ball on the game screen.
func (g *Game) DrawBall(style tcell.Style) {
	width, height := g.Screen.Size()

	g.Ball.CheckEdges(width, height)
	g.Ball.CheckCollisions(g.LocalPlayer.Paddle, g.RemotePlayer.Paddle)

	if g.Ball.HasHitLeft() {
		g.RemotePlayer.Score += 1
		pause(1000)
		g.Ball.ResetLeft(width)
	}

	if g.Ball.HasHitRight(width) {
		g.LocalPlayer.Score += 1
		pause(1000)
		g.Ball.ResetRight(width)
	}

	g.Ball.Update()

	drawSprite(g.Screen, g.Ball.X, g.Ball.Y, g.Ball.X, g.Ball.Y, style, g.Ball.Display())
}

// Draws the player scores on the game screen.
func (g *Game) DrawScores(style tcell.Style) {
	width, _ := g.Screen.Size()
	leftScore := fmt.Sprintf("← %s", strconv.Itoa(g.LocalPlayer.Score))
	rightScore := fmt.Sprintf("%s →", strconv.Itoa(g.RemotePlayer.Score))

	drawSprite(g.Screen, (width/2)-10, 1, (width/2)-7, 1, style, leftScore)
	drawSprite(g.Screen, (width/2)+10, 1, (width/2)+13, 1, style, rightScore)
}

// Checks if the game is over.
func (g *Game) GameOver() bool {
	return g.LocalPlayer.Score == g.MaxScore || g.RemotePlayer.Score == g.MaxScore
}

// Declares the winner of the game.
func (g *Game) DeclareWinner() string {
	if g.LocalPlayer.Score > g.RemotePlayer.Score {
		return "← Winner"
	} else {
		return "Winner →"
	}
}

// Listens for user input events, like keyboard input.
func (g *Game) PollEvents() {
	screen := g.Screen
	_, height := screen.Size()

	switch event := screen.PollEvent().(type) {
	case *tcell.EventResize:
		screen.Sync()
	case *tcell.EventKey:
		if isExitKey(event.Key()) {
			g.End()
		} else if event.Key() == tcell.KeyUp {
			g.LocalPlayer.Paddle.MoveUp()
		} else if event.Key() == tcell.KeyDown {
			g.LocalPlayer.Paddle.MoveDown(height)
		}
	}
}

// Draw a wait screen while waiting for a connection in server mode.
func (g *Game) DrawWaitScreen(style tcell.Style) {
	width, _ := g.Screen.Size()

	waitMessage := fmt.Sprintf("Waiting for a player to connect on port %d...", g.Port)
	drawSprite(g.Screen, (width/2)-25, 7, (width/2)+25, 7, style, waitMessage)
	g.Screen.Show()
}

// Draws the end game screen.
func (g *Game) DrawEndGame(style tcell.Style) {
	width, _ := g.Screen.Size()

	drawSprite(g.Screen, (width/2)-4, 7, (width/2)+5, 7, style, g.DeclareWinner())
	drawSprite(g.Screen, (width/2)-8, 10, (width/2)+8, 10, style, "(CTRL+C to exit)")
	g.Screen.Show()
}

// Ends the game.
func (g *Game) End() {
	g.Screen.Fini()
	if g.DebugMode {
		g.PrintErrors()
	}
	close(g.Errors)
	os.Exit(0)
}

// Prints all errors that were sent to the `Errors` channel during the game.
func (g *Game) PrintErrors() {
	for err := range g.Errors {
		fmt.Printf("Encountered error: %v\n", err)
	}
}

// Draws a sprite on the screen, a group of runes with rectangular boundaries set
// by `xStart`, `yStart`, `xEnd`, and `yStart`.
func drawSprite(s tcell.Screen, xStart, yStart, xEnd, yEnd int, style tcell.Style, text string) {
	row := yStart
	col := xStart

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++

		// If we've reached the vertical edge (last column), move down a row and
		// start from the first column.
		if col >= xEnd {
			row++
			col = xStart
		}

		// If we've reach the horizontal edge (last row), we've finished drawing
		// the sprite.
		if row > yEnd {
			break
		}
	}
}

func pause(milliseconds time.Duration) {
	time.Sleep(milliseconds * time.Millisecond)
}
