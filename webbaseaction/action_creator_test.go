package webbaseaction_test

import (
	"testing"
	"time"

	wba "github.com/kimvnhung/go_learning/webbaseaction"
	act "github.com/kimvnhung/go_learning/webbaseaction/actions"
	"github.com/stretchr/testify/require"
)

func TestActionOpenUrl(t *testing.T) {
	actions := []wba.IAction{
		&act.ActionNone{},
		&act.ActionOpenUrl{Url: "https://example.com"},
	}

	err := wba.Run(actions)
	require.NoError(t, err)
}

func TestActionSleep(t *testing.T) {
	actions := []wba.IAction{
		&act.ActionNone{},
		&act.ActionSleep{Duration: 2 * time.Second},
	}

	err := wba.Run(actions)
	require.NoError(t, err)
}

func TestRunWithCookieFile(t *testing.T) {
	actions := []wba.IAction{
		&act.ActionNone{},
		&act.ActionOpenUrl{Url: "https://facebook.com"},
		&act.ActionSleep{Duration: 2 * time.Second},
	}

	err := wba.RunWithCookiesFile(actions, "cookies.json")
	require.NoError(t, err)
}
