package main

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

var ChatGPT *openai.Client

func InitChatGPTClient(apiKey string) {
	ChatGPT = openai.NewClient(apiKey)
}

func Complete(prompt string, sysMsg string) (string, error) {
	var msg []openai.ChatCompletionMessage
	msg = append(msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: sysMsg,
	})
	msg = append(msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	resp, err := ChatGPT.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: msg,
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
