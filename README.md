netpong
=======

My attempt to implement a networked game of pong.

The game interface was inspired by [this excellent tutorial](https://earthly.dev/blog/pongo/) on using `tcell` to develop a TUI. I extended the game to support networked play between two players. It uses a streaming gRPC API for networked play.

How to play
-----------

This guide assumes you have a working `go` installation.

To install the game:

```
go install github.com/qsymmachus
```

### Hosting a game in server mode

To start a game, one player will need to start in server mode:

```
netpong --server
```

Once in the mode, the game will wait for another player to connect. By default, the game listens for connections on port 60049; you can choose a different port using the `--port` flag.

### Connecting to a game

Once a server is listening, another player can connect and start a game:

```
netpong --address localhost:60049
```

In this example we're playing against a game on the same host; use a different address and port as required.

Once connected the game will begin. The local player's paddle is on the right-hand side, and the remote player's paddle is on the left-hand side. Move your paddle up and down using the arrow key. The first player to score 5 points wins.

Development
-----------

As you're actively developing the game, you can compile and play it with this command:

```sh
go run .
```

The game uses a streaming gRPC API for networked play. Protobuf files and generated code are found in `./netpong`.

If you modify `protos/netpong.proto`, you'll need to regenerate the gRPC code with this command:

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    netpong/netpong.proto
```
