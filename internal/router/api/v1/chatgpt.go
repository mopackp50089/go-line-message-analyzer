package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	gRPCChatgptClient "github.com/onepiece010938/go-line-message-analyzer/internal/adapter/grpc/chatgpt/client"
	"github.com/onepiece010938/go-line-message-analyzer/internal/app"
	"github.com/onepiece010938/go-line-message-analyzer/internal/app/service/chatgpt"
)

type ChatGPTLineHandler struct {
	bot              *linebot.Client
	chatgpt          *chatgpt.ChatGPTService
	openaigrpcclient *gRPCChatgptClient.OpenaiGrpcClient
}

func ChatGPT(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request)
		ctx := c.Request.Context()
		events, err := app.LineBotClient.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {

				c.JSON(http.StatusBadRequest, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
			return
		}
		ChatGPTLineHandler := ChatGPTLineHandler{bot: app.LineBotClient, chatgpt: app.ChatGPTService, openaigrpcclient: app.OpenaiGrpcClient}
		for _, event := range events {
			log.Printf("Got event %v", event)
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if err := ChatGPTLineHandler.handleText(ctx, app, message, event.ReplyToken, event.Source); err != nil {
						log.Print(err)
					}

				case *linebot.FileMessage:
					if err := ChatGPTLineHandler.handleFile(message, event.ReplyToken); err != nil {
						log.Print(err)
					}
				default:
					log.Printf("Unknown message: %v", message)
				}
			case linebot.EventTypeFollow:
				if err := ChatGPTLineHandler.replyText(event.ReplyToken, "Got followed event"); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeUnfollow:
				log.Printf("Unfollowed this bot: %v", event)
			case linebot.EventTypeJoin:
				if err := ChatGPTLineHandler.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeLeave:
				log.Printf("Left: %v", event)
			case linebot.EventTypePostback:
				data := event.Postback.Data
				if data == "DATE" || data == "TIME" || data == "DATETIME" {
					data += fmt.Sprintf("(%v)", *event.Postback.Params)
				}
				if err := ChatGPTLineHandler.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeBeacon:
				if err := ChatGPTLineHandler.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown event: %v", event)
			}
		}

	}
}

func (c *ChatGPTLineHandler) handleText(ctx context.Context, app *app.Application, message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "sample":
		test, err := app.AnalyzeService.AnalyzeTest(ctx)
		if err != nil {
			return c.replyText(replyToken, "ERROR")
		}
		return c.replyText(replyToken, test)

	default:
		// --- GO version: chatGPT Model: gogpt.GPT3TextDavinci003 ---
		// log.Printf("Echo message to %s: %s", replyToken, message.Text)
		// log.Println("Call GO ChatGPT")
		// completionRequestParm := chatgpt.CompletionRequestParm{
		// 	Prompt: message.Text,
		// }
		// resp, err := c.chatgpt.CompletionRequest(ctx, completionRequestParm)
		// if err != nil {
		// 	return c.replyText(replyToken, err.Error())
		// }

		// --- PYTHON latest version: chatGPT Model: gpt-3.5-turbo
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		log.Println("Call PYTHON ChatGPT")
		err, resp := c.openaigrpcclient.Chatgpt(message.Text)
		if err != nil {
			return c.replyText(replyToken, err.Error())
		}
		return c.replyText(replyToken, resp)
	}
}

func (l *ChatGPTLineHandler) handleFile(message *linebot.FileMessage, replyToken string) error {
	return l.replyText(replyToken, fmt.Sprintf("File `%s` (%d bytes) received.", message.FileName, message.FileSize))
}

func (l *ChatGPTLineHandler) replyText(replyToken, text string) error {
	if _, err := l.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}
