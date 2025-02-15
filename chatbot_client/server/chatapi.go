package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kimvnhung/go_learning/chatbot_client/proto/protogenerated"
)

type ChatApiServer struct {
	pb.UnimplementedChatApiServer
}

func (s *ChatApiServer) GetChat(ctx context.Context, in *pb.MessageRequest) (*pb.MessageRespone, error) {
	log.Printf("Received at GetChat: %v", in.GetMessage())
	res, err := instance.GetResponse(&pb.GetResponseRequest{Message: fmt.Sprintf("Message from ChatApiServer: \"%v\"", in.GetMessage())})

	if err != nil {
		return nil, err
	}
	return &pb.MessageRespone{Message: fmt.Sprintf("Response from chatboter api: <%s>", res.GetMessage())}, nil
}
