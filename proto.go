package proto

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	DelAccServer
}

func (s *Server) Do(ctx context.Context, in *Response) (*Response, error) {
	log.Printf("Receive message body from client: %s", in.Message)
	return &Response{Message: "Hello From the Server!"}, nil
}
