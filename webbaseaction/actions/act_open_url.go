package webbaseaction

import (
	"context"
	"log"

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
	log.Println("No action to perform")
	return nil
}
