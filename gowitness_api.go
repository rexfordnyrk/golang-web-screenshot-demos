package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	getImage("www.itsfoss.com/install-docker-fedora")
	//println(fmt.Sprintf("%s-%d","https://innoverex.com", time.Now().UTC().Unix()))
}

func getImage(site string) {
	//concatenating the url string to make the request
	screenShotService := fmt.Sprintf("http://localhost:7171?url=%s%s", "https://", site)

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

	// You have to manually close the body
	//but defer closing till the method is done executing and about it exit
	defer resp.Body.Close()

	//naming file using provided URL without "/"s and current unix datetime
	filename := fmt.Sprintf("%s-%d.png", strings.Replace(site, "/", "-", -1), time.Now().UTC().Unix())

	// You can now save it to disk...
	errs := ioutil.WriteFile(filename, body, 0666)
	if errs != nil {
		log.Fatalln(errs.Error())
	}
	log.Printf("..............saved screenshot to file %s", filename)

}
