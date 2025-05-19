package main

import (
	"context"
	"google.golang.org/genai"
	"os"
)

func main() {
	ctx := context.Background()
	genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GENAI_API_KEY"),
	})
}
