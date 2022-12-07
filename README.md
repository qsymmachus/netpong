pongo
=====

My attempt to implement pong.

Based on this great project tutorial: [https://earthly.dev/blog/pongo/](https://earthly.dev/blog/pongo/)

I'm extending the project to support networked play.

Usage
-----

As you're actively developing the game, you can compile and play it with this command:

```sh
go run .
```

If you modify `protos/netpong.proto`, you'll need to regenerate the gRPC code with this command:

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    netpong/netpong.proto
```

__TODO__: Consider creating a makefile for this codegen task.
