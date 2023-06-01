package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebsiteInfo struct {
	URL   string
	Title string
}

func main() {
	var urls []string = []string{
		"https://links.hr",
		"https://instar-informatika.hr",
	}

	// channel that receives website info
	results := make(chan WebsiteInfo)

	// start goroutines to scrape websites concurrently
	for _, url := range urls {
		go func(u string) {
			wi, err := scrapeWebsite(u)
			if err != nil {
				fmt.Println(err)
				return
			}
			results <- wi
		}(url)
	}

	// wait for all goroutines to finish
	for i := 0; i < len(urls); i++ {
		info := <-results
		fmt.Printf("URL: %s\nTitle: %s\n\n", info.URL, info.Title)
	}
}

func scrapeWebsite(url string) (WebsiteInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return WebsiteInfo{}, err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return WebsiteInfo{}, err
	}

	info := WebsiteInfo{
		URL:   url,
		Title: "website title",
	}
	return info, nil
}
