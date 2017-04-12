package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//Data export
type Data struct {
	Index  string
	Status string
}

//Wait return
type Wait struct {
	DIdx string
	CIdx string
}

//Magic export
type Magic struct {
	Time string
	Dlr  Data
	Dca  Data
}

func main() {

	var allMagic Magic
	var disney, cali Data
	var rW Wait

	pDLR := make(chan string)
	pDCA := make(chan string)
	wait := make(chan Wait)

	allMagic.Time = fmt.Sprintf("%v", time.Now().Unix())
	go scrapePacked("http://www.isitpacked.com/live-crowd-trackers/disneyland/", pDLR)
	go scrapePacked("http://www.isitpacked.com/live-crowd-trackers/dca-disney-california-adventure/", pDCA)
	go scrapeMouseWait(wait)

	disney.Status = <-pDLR
	cali.Status = <-pDCA
	rW = <-wait

	disney.Index = rW.DIdx
	cali.Index = rW.CIdx

	//x, y, z := <-packed, <-packed, <-wait // receive from c
	//fmt.Println(x, y, z)

	allMagic.Dlr = disney
	allMagic.Dca = cali

	amJSON, _ := json.MarshalIndent(allMagic, "", "    ")
	dbSet(amJSON)
	dbGet()
	//fmt.Println(string(amJSON))

	//fmt.Printf("%+v\n", allMagic)
	// fmt.Println(allM.california)
	time.Sleep(180000 * time.Millisecond)
	main()

}
