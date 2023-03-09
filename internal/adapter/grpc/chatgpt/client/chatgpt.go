package openaiGrpcClient

import (
	"context"
	"fmt"
	"log"

	pb "github.com/onepiece010938/go-line-message-analyzer/internal/adapter/grpc/chatgpt/pb"
	"google.golang.org/grpc/status"
)

// func doChatgpt(c pb.OpenAiClient) {
// 	log.Println("doChatgpt was invoked")
// 	req := &pb.ChatGptRequest{
// 		Role:    "user",
// 		Content: "請告訴我吃蘋果的優點",
// 	}
// 	resp, err := c.ChatGpt(context.Background(), req)
// 	if err != nil {
// 		e, ok := status.FromError(err)
// 		if ok {
// 			log.Printf("Error message from server: %s\n", e.Message())
// 			log.Printf("Error code from server: %s\n", e.Code())
// 		} else {
// 			log.Fatalf("A non grpc error:: %v\n", err)
// 		}
// 	}
// 	log.Printf("Chatgpt reply: %s\n", resp.Response)
// }

func (openai *OpenaiGrpcClient) Chatgpt(content string) (error, string) {
	log.Println("doChatgpt was invoked")
	req := &pb.ChatGptRequest{
		Role:    "user",
		Content: content,
	}

	resp, err := openai.client.ChatGpt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			return fmt.Errorf("Error message from server: %s\n", e.Message()), ""
		} else {
			log.Fatalf("A non grpc error: %v\n", err)
			return fmt.Errorf("A non grpc error: %v\n", err), ""
		}
	}
	log.Printf("Chatgpt reply: %s\n", resp.Response)
	return nil, resp.Response

}
