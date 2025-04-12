package actions

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	wba "github.com/kimvnhung/go_learning/webbaseaction"
)

type FbPost struct {
	PostMessage string
}

type ActionGetFBPost struct {
	Url       string // The URL to open
	ClassPath string
	CallBack  func(FbPost) error
}

func (a *ActionGetFBPost) GetActionType() wba.ActionType {
	return wba.ActGetFBPost
}

func (a *ActionGetFBPost) Act(ctx context.Context) error {
	// define a variable to store element
	// var postCount int
	var nodes []*cdp.Node
	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(a.Url),
		// Wait for meaningful content
		chromedp.WaitVisible(".loaded-class", chromedp.ByQuery),
		// Optional delay
		chromedp.Sleep(2*time.Second),

		// Grab final HTML
		chromedp.OuterHTML("html", &html),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Auto red")
			// #«rvr» > span > div
			var text string
			err := chromedp.InnerHTML(a.ClassPath, &text).Do(ctx)
			log.Printf("Text: %s\n", text)
			return err
		}),
	)
	log.Printf("Post count: %d\n", len(nodes))
	log.Printf("Opened URL: %s, ClassPath: %s\n", a.Url, a.ClassPath)
	return err
}
