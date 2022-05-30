// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	getChromedpScreenShot("www.itsfoss.com/install-docker-fedora",100)
}

func getChromedpScreenShot(site string, quality int) {
	//forming url to be captured
	screenShotUrl := fmt.Sprintf("https://%s/", site)

	//byte slice to hold captured images in bytes
	var buf []byte

	//setting image file extension to png but
	var ext string = "png"
	//if image quality is less than 100 file extension is jpeg
	if quality < 100 {
		ext = "jpeg"
	}

	log.Printf("................making request for screenshot using %s", screenShotUrl)

	//setting options for headless chrome to execute with
	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1400, 900))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	//setup context with options
	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)

	defer acancel()

	// create context
	ctx, cancel := chromedp.NewContext( actx,
	)
	defer cancel()


	tasks:= chromedp.Tasks{
		//loads page of the URL
		chromedp.Navigate(screenShotUrl),

		//waits for 5 secs
		chromedp.Sleep(5*time.Second),

		//Captures Screenshot with current window size
		chromedp.CaptureScreenshot(&buf),

		//capture full-page screenshot
		//chromedp.FullScreenshot(&buf,quality),
	}
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}

	//naming file using provided URL without "/"s and current unix datetime
	filename := fmt.Sprintf("%s-%d.%s",strings.Replace(site,"/","-",-1), time.Now().UTC().Unix(), ext)

	//write byte slice data of standard screenshot to file
	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		log.Fatal(err)
	}

	//log completion and file name to
	log.Printf("..............saved screenshot to file %s", filename)
}