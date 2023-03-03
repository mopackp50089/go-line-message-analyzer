package chatgpt

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type CompletionRequestParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
	Prompt string
}

func (i *ChatGPTService) CompletionRequest(ctx context.Context, parm CompletionRequestParm) (string, error) {
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   256,
		Prompt:      parm.Prompt,
		Temperature: 0.7,
	}

	resp, err := i.ChatGPTServiceClient.CreateCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("fail to CreateCompletion: %v", err)
	}
	fmt.Println("chatGPT return!!")
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text, nil
}
