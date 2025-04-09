package webbaseaction_test

import (
	"testing"

	wba "github.com/kimvnhung/go_learning/webbaseaction"
	act "github.com/kimvnhung/go_learning/webbaseaction/actions"
	"github.com/stretchr/testify/require"
)

func TestNewAction(t *testing.T) {
	actions := []wba.IAction{
		&act.ActionNone{},
		&act.ActionOpenUrl{Url: "https://example.com"},
	}

	err := wba.Run(actions)
	require.NoError(t, err)
}
