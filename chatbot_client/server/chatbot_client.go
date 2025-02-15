package server

import (
	"context"
	"log"
	"sync"

	pb "github.com/kimvnhung/go_learning/chatbot_client/proto/protogenerated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatbotClient struct {
	cancel context.CancelFunc
	client pb.ChatboterClient
	ctx    context.Context
}

var (
	instance *ChatbotClient
	once     sync.Once
)

func init() {
	log.Println("Starting Chatbot Client Internal")
	NewChatbotClient()
}

func NewChatbotClient() {
	once.Do(func() {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		instance = &ChatbotClient{}

		instance.client = pb.NewChatboterClient(conn)
		instance.ctx, instance.cancel = context.WithCancel(context.Background())
	})
}

func (c *ChatbotClient) GetResponse(req *pb.GetResponseRequest) (*pb.GetResponseResponse, error) {
	log.Printf("Sending request to Chatbot Service: %v", req)
	return instance.client.GetResponse(instance.ctx, req)
}
