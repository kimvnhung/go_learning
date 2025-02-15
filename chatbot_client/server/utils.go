package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/kimvnhung/go_learning/chatbot_client/proto/protogenerated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartServices(apiOutPort int, apiServicePort int, chatbotServicePort int) error {
	// Start the apiServer
	err := establishApiService(apiServicePort)
	if err != nil {
		return err
	}

	// Connect to the Chatbot Service
	// err = connectToChatbotService(chatbotServicePort)
	// if err != nil {
	// 	log.Printf("Error connecting to Chatbot Service: %v", err)
	// }

	// Start the apiOutServer
	return establishApiOut(apiOutPort, apiServicePort)
}

func establishApiService(port int) error {
	// Start the API Service
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterChatApiServer(s, &ChatApiServer{})
	// Serve gRPC server
	log.Printf("Api internal service listening at %v", lis.Addr())
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	return nil
}

// func connectToChatbotService(port int) error {
// 	// Connect to the Chatbot Service
// 	return errors.New("not implemented")
// }

func establishApiOut(outPort int, apiServicePort int) error {
	// Start the API Out Service
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%d", apiServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	// Register Greeter

	err = pb.RegisterChatApiHandler(ctx, gwmux, conn)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", outPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%d\n", outPort)
	log.Fatalln(gwServer.ListenAndServe())

	return nil
}
