package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Article struct {
	Title string
	Date string
	URL string
	Publisher string
	Synopsis string
	Content string
}

func scrapeHindustanTimes() (articles []Article) {
	//To store the list of articles scraped
	articlesList := []Article{}
	//Instantiate new Colly Collector 
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request){
		log.Println("Visiting ", r.URL.String())
	})
	//First scrape the article links on the page and then visit the page one by one
	c.OnHTML("body > section:nth-child(20) > div > div.row > div.col-xl-6.col-lg-4.col-md-12.col-sm-12.col-12.result h3 a[href]", func(e *colly.HTMLElement){
			link := e.Attr("href")
			if strings.Contains(link, "-ai") || strings.Contains(link, "ai-"){
				e.Request.Visit(link)
			}
			
	})
	//Then scrape the page for article title, date, synopsis, url, content etc
	c.OnHTML("body > section.mt-4 > div > div > div.col-xl-9.col-lg-8.col-md-12.col-sm-12.col-12.storyline", func(e *colly.HTMLElement){
			article := Article{}
			article.Title = e.ChildText("h1")
			article.Synopsis = e.ChildText("h2")
			article.Date = e.ChildText("div.update-publish-time > p > span")
			article.URL = e.Request.URL.String()

			e.DOM.Find("div.col-xl-9.col-lg-8.col-md-12.col-sm-12.col-12.storyline p").Each(func(_ int, s *goquery.Selection){
				article.Content = article.Content + s.Text()

			})
			relatedTopics := e.ChildText("p.related-topics-list")
			article.Content = strings.Replace(article.Content, relatedTopics,"", -1)
			articlesList = append(articlesList, article)
	})
	c.Visit("https://thehindu.com/sci-tech/technology/")
	c.Wait()
	fmt.Println("List of Scraped Articles -->", articlesList)
	return articlesList
}
