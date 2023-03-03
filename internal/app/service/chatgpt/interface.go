package chatgpt

type ChatGPTClient interface {
	GetMessageCache(input string) string
}
