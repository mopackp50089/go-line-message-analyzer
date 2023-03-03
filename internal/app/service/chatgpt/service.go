package chatgpt

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatGPTService struct {
	ChatGPTServiceClient *gogpt.Client
}

type ChatGPTServiceParam struct {
	OpenApiKey string
}

func NewChatGPTService(_ context.Context, param ChatGPTServiceParam) *ChatGPTService {
	gpt := gogpt.NewClient(param.OpenApiKey)
	return &ChatGPTService{
		ChatGPTServiceClient: gpt,
	}
}
