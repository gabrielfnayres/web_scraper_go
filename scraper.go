package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Products struct{
  url, image, name, price string
}

func main(){

  fmt.Print("oi")
  c := colly.NewCollector()


    // defining some callbacks
  c.OnRequest(func(r *colly.Request){
    fmt.Println("Visiting: ", r.URL)
  })
  // error callback
  c.OnError(func(r *colly.Response, err error){
    log.Println("Something went wrong: ", err)
  })
  // default response
  c.OnResponse(func(r *colly.Response){
    fmt.Println("Page visited: ", r.Request.URL)
  })
 // printing all pages 
  c.OnHTML("a", func(e *colly.HTMLElement){
    fmt.Println( e.Attr("href"))
  })

  // verifying if it got scraped or not

  c.OnScraped(func(r *colly.Response){
    fmt.Println(r.Request.URL, " done!")
  })

  c.Visit("https://scrapeme.live/shop/") 
}
