package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	getUrlBoxImage("www.innoverex.com", "Your-API-Key", "png")
	//println(fmt.Sprintf("%s-%d","https://innoverex.com", time.Now().UTC().Unix()))
}

func getUrlBoxImage(site string, apiKey string, format string) {
	//concatenating the api key and document format to make the request
	screenShotService := fmt.Sprintf("https://api.urlbox.io/v1/%s/%s?", apiKey, format)

	//creating a map of key-value pairs of Urlbox API options
	params := url.Values{
		"url":                 {site},
		"format":              {"png"},
		"width":               {"1400"},
		"height":              {"690"},
		"full_page":           {"true"},
		"block_ads":           {"true"},
		"hide_cookie_banners": {"true"},
		"click_accept":        {"true"},
	}

	//Configuring the request with the method, URL and body
	req, err := http.NewRequest("GET", screenShotService, nil)

	//encode values into URL encoded form/query parameters
	req.URL.RawQuery = params.Encode()

	//printing out to console the entire request url with params. You can comment this out
	fmt.Println(req.URL.String())

	//Create a default HTTP client to make the request
	client := &http.Client{}

	//making the get request to the Urlbox screenshot API
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	//defer closing of body till the method is done executing and about it exit
	defer resp.Body.Close()

	//We Read the response body (the image) on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//naming file using provided URL without "/"s and current unix datetime
	filename := fmt.Sprintf("%s-%d.png", strings.Replace(site, "/", "-", -1), time.Now().UTC().Unix())

	// You can now save it to disk...
	errs := ioutil.WriteFile(filename, body, 0666)
	if errs != nil {
		log.Fatalln(errs.Error())
	}

	log.Printf("..............saved screenshot to file %s", filename)
}
