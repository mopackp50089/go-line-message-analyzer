package analyze

import (
	"context"
	"fmt"
)

type CreateAnalyzeParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
}

type AnalyzeMessageParm struct {
}

func (i *AnalyzeService) CreateAnalyze(ctx context.Context, param CreateAnalyzeParm) error {
	fmt.Println("AnalyzeService->func CreateAnalyze()")
	// var result message.MessageDomainResult
	// result = message.MessageDomainFunc("aabcccc")
	// fmt.Println(result)
	i.analyzeServiceCache.GetMessageCache("")

	return nil
}

func (i *AnalyzeService) AnalyzeTest(ctx context.Context) (string, error) {
	fmt.Println("AnalyzeService->func AnalyzeTest()")
	// var result message.MessageDomainResult
	// result = message.MessageDomainFunc("aabcccc")
	// fmt.Println(result)
	return "TESTTESTTEST", nil
}

// func (i *AnalyzeService) StartAnalyze(ctx context.Context) {
// 	var resultMessage []string
// 	var resultCall []string
// 	var resultPhoto []string
// 	var resultSticker []string
// 	for scanner.Scan() {
// 		sender, timer, content, class := DomainAnalyzeCalssfi(scanner.Text())
// 		switch class {
// 		case "message":
// 			{
// 				DomainAnalyzeMessage(sender, timer, content, &resultMessage)
// 			}
// 		case "call":
// 			{
// 				DomainAnalyzeCall(sender, timer, content, &resultCall)
// 			}
// 		case "photo":
// 			{
// 				DomainAnalyzePhoto(sender, timer, content, &resultPhoto)
// 			}
// 		case "sticker":
// 			{
// 				DomainAnalyzeSticker(sender, timer, content, &resultSticker)
// 			}
// 		}
// 	}
// 	DomainMessage(&resultMessage)

// }

// // Message Domain
// func DomainMessage(resultMessage *[]string) {
// 	domainMessageSetCache()
// 	domainMessageGetCache()
// 	domainMessageDraw()
// }

// func DomainAnalyzeCalssfi(row string) {
// 	var class string
// 	line := Strings.Split(row, "\t")
// 	sender, timer, content := line[0], line[1], line[2]
// 	if content == "message" {
// 		class = "message"
// 	} else if content == "call" {
// 		class = "call"
// 	} else if content == "photo" {
// 		class = "photo"
// 	} else {
// 		class = "sticker"
// 	}
// 	return sender, timer, content, class
// }

// // Analyze Domain
// func DomainAnalyzeMessage() {

// }

// func DomainAnalyzeCall() {

// }

// func DomainAnalyzePhoto() {

// }

// func DomainAnalyzeSticker() {

// }
