package grpc

import (
	"context"
	"log"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/okzmo/machin/fbs/Greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

// JUST FOR TESTING

type server struct {
	Greeter.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *Greeter.HelloRequest) (*flatbuffers.Builder, error) {
	v := request.Name()
	var m string
	if v == nil {
		m = "Unknown"
	} else {
		m = string(v)
	}
	b := flatbuffers.NewBuilder(0)
	idx := b.CreateString("What's up " + m)
	Greeter.HelloReplyStart(b)
	Greeter.HelloReplyAddMessage(b, idx)
	b.Finish(Greeter.HelloReplyEnd(b))
	return b, nil
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	encoding.RegisterCodec(flatbuffers.FlatbuffersCodec{})
	Greeter.RegisterGreeterServer(s, &server{})

	log.Println("Server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
