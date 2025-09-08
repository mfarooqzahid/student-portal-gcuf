package ai

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

type GenAI struct {
	Client *genai.Client
}

func NewGenAiClient() GenAI {

	apikey := os.Getenv("OPENAI_API_KEY")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apikey})
	if err != nil {
		log.Fatal(err)
	}

	return GenAI{
		Client: client,
	}
}

func (g *GenAI) Generate(ctx context.Context, config *genai.GenerateContentConfig, instruction string) (string, error) {
	resp, err := g.Client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(instruction), config)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
