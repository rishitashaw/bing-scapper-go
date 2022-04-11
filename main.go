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

func buildBingURLs(searchTerm, country string, pages, count int) ([]string, error) {
toScrape:=[]string{}
searchTerm=strings.Trim(searchTerm," ")
searchTerm=strings.Replace(searchTerm," ","+",-1)
if countryCode, found:= bingDomains[country]; found{
	for i:=0;i<pages;i++{
		first:=firstParameter(i,count)
		scarpeURL:=fmt.Sprintf("http://bing.com/search?q=%s&first=%d&count=%d%s",searchTerm,first,count, countryCode)
		toScrape=append(toScrape, scarpeURL)
	}
}else{
	fmt.Errorf("country(%s) is not found",country)
	return nil, err
}
return toScrape, nil
}

func firstParameter(number, count int) int {
	if number==0{
		return number+1
	}
	return number*count+1
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
	return results, nil
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