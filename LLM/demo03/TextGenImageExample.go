package main

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI})

	if err != nil {
		fmt.Println("create client failed:", err)
		return
	}

	result, err := client.Models.GenerateContent(ctx,
		"gemini-2.0-flash-preview-image-generation",
		genai.Text("Generate a picture of three matches dancing. The matches look like people"),
		&genai.GenerateContentConfig{
			ResponseModalities: []string{"TEXT", "IMAGE"},
		},
	)
	if err != nil {
		fmt.Println("generate content failed:", err)
		return
	}
	if result.Candidates != nil {
		for _, content := range result.Candidates[0].Content.Parts {
			if content.Text != "" {
				fmt.Println("Text:", content.Text)
			} else if content.InlineData != nil {
				currentDir, err := os.Getwd()
				fmt.Println("Current Dir:", currentDir)
				if err != nil {
					fmt.Println("get current dir failed:", err)
				}
				err = os.WriteFile("./LLM/demo03/image.jpg", content.InlineData.Data, 0644)
				if err != nil {
					fmt.Println("write image failed:", err)
				}
				fmt.Println("Image saved to image.jpg")
			}
		}
	}
}
