**This is Scraper code for Hindustantimes.com/sci-tech/technology/**

The article links are first scraped on the main page mentioned above and then each link is visite separately, 
To scrape the article links we need to use the appropriate CSS selectors

**CSS Selectors for Article link**

   -  body > section:nth-child(19) > div > div.row > div.col-xl-6.col-lg-4.col-md-12.col-sm-12.col-12.result > div:nth-child(2) > div > h3 > a
   -  However we use only "body > section:nth-child(19) > div > div.row > div.col-xl-6.col-lg-4.col-md-12.col-sm-12.col-12.result  h3  a[href]"


each article link is visited and the article's details like Title, Date, Synopsis, Content etc is scraped

**CSS Selectors for Article details**

body > section.mt-4 > div > div > div.col-xl-9.col-lg-8.col-md-12.col-sm-12.col-12.storyline 

 the above selector houses the entire article with title, date, synopisis, content + related topics link in the end 

We then select the child elements by using e.ChildText("h1") ---> for Title
                                           e.ChildText("h2") ---> for synopsis
                                           e.ChildText("div.update-publish-time > p > span") ---> for date time

We then use goquery package to navigate the DOM and loop through the p elements that contain the Content

