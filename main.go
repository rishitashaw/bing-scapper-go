package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

bingDomains := map[string]string{
	"com":""
}

type SearchResult struct {
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",

}

func randomUserAgent() string {

}

func buildBingURLs() string {

}

func scrapeClientRequest(){

}

func BingScrape(searchTerm string, country string, pages int, count int, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}

	bingPages, err := buildBingURLs(searchTerm, country, pages, count)
	
	if err != nil {
		return nil, err
	}
	
	for _, page := range bingPages {
		rank:=len(results)
		res,err:=scrapeClientRequest(page)
		if err !=nil{
			return nil, err
		}
		data,err:=bingResultParser(res,rank)
		if err != nil {
			return nil, err
		}
		for _,result := range data{
			results=append(results,result)
		}
	time.Sleep(time.Duration(backoff)*time.Second)
	}
}

func bingResultParser(){

}

func main() {
	res,err := BingScrape("rishita shaw", "com",2,30,30)
	if err == nil {
		for _,res := range res{
			fmt.Println(res)
		}
	}else{
		fmt.Println(err)
	}
}