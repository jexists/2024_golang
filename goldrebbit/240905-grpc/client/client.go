package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

var id = flag.String("id", "unknow", "the id name")
var serverAddr = flag.String("addr", "localhost:50051", "the")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials)
	if err != nil {
		log.Fatal("fail to dia %v", err)
	}
	defer conn.Close()
	client := pb.New
	runChat(client)
}

func runChat(client pb.ChatServiceClient) {
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatal("client.Chat fail %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatal("client.Chat fail %v", err)
			}
			log.Printf("sender: %s message %s", in.Sender, in.Message)
		}
	}()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		msg := s.Text()
		if strings.ToLower(msg) == "exit" {
			break
		}
		stream.Send(&pb) ???
	}
}
