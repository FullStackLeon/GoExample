package main

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"os"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("GEMINI_API_KEY environment variable not set")
		return
	}
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("Explain how AI works in a few words"),
		nil,
	)

	fmt.Println(result)
}
