package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client := cohereclient.NewClient(cohereclient.WithToken(os.Getenv("COHERE_API_KEY")))

	documents := []*cohere.V2ChatRequestDocumentsItem{
		{
			Document: &cohere.Document{
				Data: map[string]interface{}{
					"title": "会社概要",
					"text":  "私たちはAIを活用したソフトウェア開発を行う企業です。2020年に設立されました。",
				},
			},
		},
		{
			Document: &cohere.Document{
				Data: map[string]interface{}{
					"title": "サービス",
					"text":  "主なサービスはチャットボット開発、データ分析、AI導入支援です。",
				},
			},
		},
	}

	response, err := client.V2.Chat(context.Background(), &cohere.V2ChatRequest{
		Model: "command-a-03-2025",
		Messages: cohere.ChatMessages{
			{
				Role: "system",
				System: &cohere.SystemMessageV2{
					Content: &cohere.SystemMessageV2Content{
						String: "与えられた文書に基づいて回答してください。不明な点は推測せず不明と答えてください。",
					},
				},
			},
			{
				Role: "user",
				User: &cohere.UserMessageV2{
					Content: &cohere.UserMessageV2Content{
						String: "この会社はどんなサービスを提供していますか？",
					},
				},
			},
		},
		Documents: documents,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 本文表示
	for _, block := range response.Message.Content {
		if block.Text != nil {
			fmt.Println(block.Text.Text)
		}
	}

	// citation 表示
	for _, citation := range response.Message.Citations {
		fmt.Printf("%+v\n", citation)
	}
}
