package actions

import (
	"context"
	"log"
	"time"

	wba "github.com/kimvnhung/go_learning/webbaseaction"
)

type ActionSleep struct {
	Duration time.Duration // Duration to sleep
}

func (a *ActionSleep) GetActionType() wba.ActionType {
	return wba.ActNone
}
func (a *ActionSleep) Act(ctx context.Context) error {
	// Do nothing
	log.Printf("Sleeping for %v seconds", a.Duration.Seconds())
	time.Sleep(a.Duration)
	log.Println("Awake now")
	return nil
}
