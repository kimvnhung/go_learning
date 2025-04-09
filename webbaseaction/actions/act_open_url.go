package actions

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	wba "github.com/kimvnhung/go_learning/webbaseaction"
)

type ActionOpenUrl struct {
	Url string // The URL to open
}

func (a *ActionOpenUrl) GetActionType() wba.ActionType {
	return wba.ActOpenURL
}
func (a *ActionOpenUrl) Act(ctx context.Context) error {
	// Do nothing
	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate(a.Url),
		chromedp.Title(&title),
	)

	log.Printf("Opened URL: %s, Title: %s\n", a.Url, title)
	return err
}
