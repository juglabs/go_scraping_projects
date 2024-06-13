package main

import (
	"context"
	"fmt"
	"log"
//	"os"		
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
//	"time"
)

func main() {
	ctx := context.Background()
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		fmt.Println("Error instantiating LLM - ", err)
		log.Fatal(err)
	}

	articlesList := scrapeHindustanTimes()

	query := "Summarize this News article - "
	articleSummary := ""
	var response string
	for _, art := range articlesList {
		query = query + art.Content
		response, err = llms.GenerateFromSinglePrompt(ctx, llm, query)
		if err != nil {
			fmt.Printf("Error generating LLM response for Article %s - %s", art.Title, err)
		}
		articleSummary = articleSummary + response
	}

	fmt.Println("Articles Summary --->", articleSummary)

}
