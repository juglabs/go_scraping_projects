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
	//Instantiating a new llama2 LLM with ollama
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		fmt.Println("Error instantiating LLM - ", err)
		log.Fatal(err)
	}
	//Get the articles from the scraper function defined in scraper_et.go
	//actually it should be ht because its for Hindustantimes.com/sci-tech/technology
	articlesList := scrapeHindustanTimes()

	query := "Summarize this News article - "
	articleSummary := ""
	var response string
	//Here we loop over each article and generate their summary using the LLM and then combine the summary
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
