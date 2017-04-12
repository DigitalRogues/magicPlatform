package main

import (
	//	"fmt"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func scrapeMouseWait(c chan Wait) {
	fmt.Printf("start wait, %v \n", time.Now())
	doc, err := goquery.NewDocument("http://www.mousewait.com/disneyland/")
	if err != nil {
		log.Fatal(err)
	}
	headings := doc.Find(".parkheading")
	var wait Wait
	wait.DIdx = splitIndex(headings.Eq(1).Text())
	wait.CIdx = splitIndex(headings.Eq(2).Text())

	c <- wait
}

func scrapePacked(url string, c chan string) {
	fmt.Printf("start packed, %v \n", time.Now())
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	//var status string
	var status string
	doc.Find(".entry-content > div:first-of-type ").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i2 int, s2 *goquery.Selection) {
			id, _ := s2.Attr("id")
			if id != "button-0" {
				status = s2.Text()
			}
		})
	})
	c <- status
}
