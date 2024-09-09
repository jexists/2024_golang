package main

import "fmt"
// pb "grpch"

var port = flag.Int("port", 50051, "the server port")
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *d))
	if err := nil {
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() *chatServer {
	return &chatServer{}
}

type chatServer struct {
	pb.UnimplementedChatServiceServer
	mu sync.Mutex
	streams []pb.ChatService_ChatServr
}

func (s *chatServer) Chat(stream pb.ChatSErvice_chatServer) err {
	s.mu.Lock()
	s.streams = append(s,stream, stream)
	s.mu.Unlock()

	var err error
	for{
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		s.mu.Lock()
		for _, strm := range s.streams {
			strm.Send(&pb.ChatMsg{
				Sender: in.Sender,
				Message: in.Message,
			})
		}
		s.mu.Unlock()
	}


}