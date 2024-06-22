package main

import (
//	"fmt"
	"log"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
)

//Setting the Database and Table for the news links to be fetched
var Database = "newsarticles"
var Table = "nbcnews"

type Article struct {
	Title string
	Date string
	URL string
	Synopsis string
	Content string
}

func linksContain(links []string, link string) bool {
	for _, l := range links{
		if l == link {
			return true
		}
	}
	return false
}

func main() {
	c := colly.NewCollector()
	links := []string{}
	articleList := []Article{}

	//Get the already scraped links from the Database 
	//and scrape only the new links 
	links, err := getLinksFromDB(Database, Table)
	if err != nil {
		log.Println("Failed to read the DATABASE for article links", err)
		panic(err.Error())
	}

	c.OnRequest(func(r *colly.Request){
		log.Println("Visiting Initial URL -> ", r.URL.String())
	})

	c.OnHTML("div.styles_left__YWGcJ.layout-grid-item.grid-col-8-l > section > div > div.styles_itemsContainer__saJYW div.wide-tease-item__info-wrapper.flex-grow-1-m a[href]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		if !linksContain(links, link){
			links = append(links, link)
			log.Println(link)
			e.Request.Visit(link)
		}	
	})

	c.OnHTML("#content > div:nth-child(7) > div > article", func(e *colly.HTMLElement){
		article := Article{}
		e.DOM.Find("section > div.article-hero__bg-container > header > div > h1").Each(func(i int, s *goquery.Selection){
			article.Title = s.Text()		
		})
				
		e.DOM.Find("section > div.article-hero__bg-container > header > div > div").Each(func(i int, s *goquery.Selection){
			article.Synopsis = s.Text()		
		})

		e.DOM.Find("div.article-body > div > div.article-body__section.layout-grid-container.article-body__last-section.article-body__first-section > div.article-body.layout-grid-item.layout-grid-item--with-gutter-s-only.grid-col-10-m.grid-col-push-1-m.grid-col-6-xl.grid-col-push-2-xl.article-body--custom-column > section > div.article-body__date-source > time").Each(func(i int, s *goquery.Selection){
			article.Date = s.Text()		
		})	

		
		e.DOM.Find("div.article-body > div > div.article-body__section.layout-grid-container.article-body__last-section.article-body__first-section > div.article-body.layout-grid-item.layout-grid-item--with-gutter-s-only.grid-col-10-m.grid-col-push-1-m.grid-col-6-xl.grid-col-push-2-xl.article-body--custom-column p").Each(func(i int, s *goquery.Selection){
			article.Content = article.Content + s.Text()
		})

		article.URL = e.Request.URL.String()
		articleList = append(articleList, article)
	})


	c.Visit("https://www.nbcnews.com/artificial-intelligence")
	c.Wait()
	writeToTable(articleList)
	displayTableContent()

//	log.Println("Scraped articles --->", articleList)

}
