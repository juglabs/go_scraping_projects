package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

type Country struct {
	Country    string
	Capital    string
	Population int
	AreaSqkm   float64
}

func main() {
	// Initialize the collector
	c := colly.NewCollector()

	countriesList := []Country{}

	// Extract the country information
	c.OnHTML("div.col-md-4", func(e *colly.HTMLElement) {
		countrySingle := Country{
			Country: e.ChildText("h3"),
			Capital: e.ChildText("span.country-capital"),
		}
		
		populationStr := e.ChildText("span.country-population")
		population, err1 := strconv.Atoi(populationStr)
		if err1 != nil {
			fmt.Println("Error converting population:", err1)
		} else {
			countrySingle.Population = population
		}

		areaStr := e.ChildText("span.country-area")
		area, err2 := strconv.ParseFloat(areaStr, 64)
		if err2 != nil {
			fmt.Println("Error converting area:", err2)
		} else {
			countrySingle.AreaSqkm = area
		}

		countriesList = append(countriesList, countrySingle)
		fmt.Printf("%+v\n", countrySingle)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Initial visit to the website
	err := c.Visit("https://scrapethissite.com/pages/simple/")
	if err != nil {
		fmt.Printf("Failed to visit the initial URL: %v\n", err)
	}

	// Print all collected countries
	fmt.Println("Collected countries:", countriesList)
}

