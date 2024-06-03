package main

//Including the standard package fmt
//goquery package is required for DOM access
import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"os"
)

//Simple data structure to store individual Article's details. Nothing fancy
type article struct {
	Title string
	URL string
	Content string
}

func main() {
	//Instantiate default collector
	c := colly.NewCollector()
	//Instantiating slice of article to store multiple articles
	articleList := []article{}
	//Storing the scraped artilces in a .txt file 
	file, err := os.Create("indianexpress-articles.txt")
	if err!= nil {
		fmt.Println("Error creating .txt file", err)
	}
	defer file.Close()	
	//Callback to access individual article's link which will be visited
	c.OnHTML("ul.article-list a[href][title]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		e.Request.Visit(link)
	})
	//Callback to scrape individual article's details - Title, URL & Content
	c.OnHTML("div.ie_single_story_container", func(e *colly.HTMLElement){
		articleSingle := article{}
		articleSingle.Title = e.ChildText("h1")
		articleSingle.URL = e.Request.URL.String()
	
		e.DOM.Find("div.story_details p").Each(func(_ int, s *goquery.Selection){
			articleSingle.Content = articleSingle.Content + s.Text()
		})
		articleList = append(articleList, articleSingle)
	})

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting initial URL ", r.URL.String())
	})

	
	c.Visit("https://indianexpress.com/article/technology/")
	//Wait till all the visit requests are completed
	c.Wait()	
	fmt.Println("All article details:")
	for _, a := range articleList {
		fmt.Printf("Title: %s\nURL: %s\nContent: %s\n\n", a.Title, a.URL, a.Content)
	}
	//Json encode the .txt file before writing the articles
	encoder := json.NewEncoder(file)
	//iterate over the articles on write to the encoded .txt file
	for _, art := range articleList {
		if err:= encoder.Encode(art); err!= nil {
			fmt.Println("Error encoding article:", err)
		}
	}
}
