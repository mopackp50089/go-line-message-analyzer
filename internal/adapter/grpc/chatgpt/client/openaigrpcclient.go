package openaiGrpcClient

import (
	pb "github.com/onepiece010938/go-line-message-analyzer/internal/adapter/grpc/chatgpt/pb"
	"google.golang.org/grpc"
)

// var addr string = "192.168.2.39:7727"

// func main() {
// 	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("fail to connect: %v\n", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewOpenAiClient(conn)
// 	doChatgpt(client)
// }

type OpenaiGrpcClient struct {
	client pb.OpenAiClient
}

func NewOpenaiGrpcClient(conn *grpc.ClientConn) *OpenaiGrpcClient {
	client := pb.NewOpenAiClient(conn)
	return &OpenaiGrpcClient{
		client: client,
	}
}
