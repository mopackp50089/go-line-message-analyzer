package app

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/onepiece010938/go-line-message-analyzer/internal/adapter/cache"
	gRPCChatgptClient "github.com/onepiece010938/go-line-message-analyzer/internal/adapter/grpc/chatgpt/client"
	serviceAnalyze "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/analyze"
	serviceChatgpt "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/chatgpt"
	serviceMessage "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/message"
)

type Application struct {
	// JobService   *serviceJob.JobService
	// ImageService *serviceImage.ImageService
	OpenaiGrpcClient *gRPCChatgptClient.OpenaiGrpcClient
	AnalyzeService   *serviceAnalyze.AnalyzeService
	MessageService   *serviceMessage.MessageService
	LineBotClient    *linebot.Client
	ChatGPTService   *serviceChatgpt.ChatGPTService
}

func NewApplication(ctx context.Context, cache cache.CacheI, lineBotClient *linebot.Client, OpenApiKey string, openaiGrpcClient *gRPCChatgptClient.OpenaiGrpcClient) *Application {

	// Create application
	app := &Application{
		OpenaiGrpcClient: openaiGrpcClient,
		ChatGPTService: serviceChatgpt.NewChatGPTService(ctx, serviceChatgpt.ChatGPTServiceParam{
			OpenApiKey: OpenApiKey,
		}),
		LineBotClient: lineBotClient,
		MessageService: serviceMessage.NewMessageService(ctx, serviceMessage.MessageServiceParam{
			MessageServiceCache: cache,
		}),
		AnalyzeService: serviceAnalyze.NewAnalyzeService(ctx, serviceAnalyze.AnalyzeServiceParam{
			AnalyzeServiceCache: cache,
		}),
	}
	return app
}
