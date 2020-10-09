package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

const url string = "https://www.washingtonpost.com/news-sitemaps/index.xml"

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	var s SitemapIndex
	var n News

	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	newsMap := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, _ := range newsMap {
		fmt.Println(idx)
	}
}