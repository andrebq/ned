package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/andrebq/ned/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type (
	editorServer struct{}
	bufferServer struct{}
)

func (es *editorServer) GetBuffers(ctx context.Context, q *api.BufferQuery) (*api.BufferList, error) {
	return nil, errors.New("not implemented")
}

func (bs *bufferServer) GetContent(ctx context.Context, b *api.BufferIdentity) (*api.LineList, error) {
	return nil, errors.New("not implemented")
}

func (bs *bufferServer) WatchLines(b *api.BufferIdentity, srv api.Buffers_WatchLinesServer) error {
	tick := time.NewTicker(time.Second * 1)
	for {
		now := <-tick.C
		line := api.Line{
			Id:       0,
			Contents: now.Format(time.RFC3339),
			Number:   1,
		}
		err := srv.Send(&line)
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 18080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterEditorServer(s, &editorServer{})
	api.RegisterBuffersServer(s, &bufferServer{})

	s.Serve(lis)
}
