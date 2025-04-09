package webbaseaction

import (
	"context"

	"github.com/chromedp/chromedp"
)

type ActionType int

const (
	ActNone    ActionType = iota // Do nothing
	ActOpenURL                   // Open a URL in the default browser
)

type IAction interface {
	GetActionType() ActionType     // Get the type of action
	Act(ctx context.Context) error // Perform the action
}

func Run(actions []IAction) error {
	// Create a parent context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	for _, action := range actions {
		if err := action.Act(ctx); err != nil {
			return err
		}
	}
	return nil
}
