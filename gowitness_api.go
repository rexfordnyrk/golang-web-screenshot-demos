package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	getImage("www.innoverex.com")
}

func getImage(site string){
	//concatenating the url string to make the request
	screenShotService := fmt.Sprintf("http://localhost:7171?url=%s%s","https://", url.QueryEscape(site))

	log.Printf("................making request for screenshot using %s", screenShotService)
	//makng the get request to the gowitness screenshot service
	resp, err := http.Get(screenShotService)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body (the image )on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// You have to manually close the body.
	defer resp.Body.Close()

	//naming file using domain name and current unix datetime
	filename := fmt.Sprintf("%s-%d.png",site, time.Now().UTC().Unix())


	// You can now save it to disk...
	errs := ioutil.WriteFile(filename, body, 0666)
	if errs != nil {
		log.Fatalln(errs.Error())
	}
	log.Printf("................saved screenshot to file %s", filename)
}