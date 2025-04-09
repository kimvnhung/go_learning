package webbaseaction

import (
	"context"
	"log"

	wba "github.com/kimvnhung/go_learning/webbaseaction"
)

type ActionNone struct {
}

func (a *ActionNone) GetActionType() wba.ActionType {
	return wba.ActNone
}
func (a *ActionNone) Act(ctx context.Context) error {
	// Do nothing
	log.Println("No action to perform")
	return nil
}
