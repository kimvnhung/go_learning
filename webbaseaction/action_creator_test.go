package webbaseaction_test

import (
	"fmt"
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
	domain := "www.facebook.com"
	actions := []wba.IAction{
		&act.ActionNone{},
		&act.ActionOpenUrl{Url: fmt.Sprintf("https://%s", domain)},
		&act.ActionSleep{Duration: 2 * time.Second},
	}

	err := wba.RunWithCookiesFile(actions, fmt.Sprintf("%s_cookies.json", domain))
	require.NoError(t, err)
}

func TestGetFbPost(t *testing.T) {
	domain := "www.facebook.com"
	actions := []wba.IAction{
		&act.ActionGetFBPost{
			Url:       fmt.Sprintf("https://%s/me", domain),
			ClassPath: "#mount_0_0_lN > div > div:nth-child(1) > div > div.x9f619.x1n2onr6.x1ja2u2z > div > div > div.x78zum5.xdt5ytf.x1t2pt76.x1n2onr6.x1ja2u2z.x10cihs4 > div.x78zum5.xdt5ytf.x1t2pt76 > div > div > div.x6s0dn4.x78zum5.xdt5ytf.x193iq5w > div.x9f619.x193iq5w.x1talbiv.x1sltb1f.x3fxtfs.x1swvt13.x1pi30zi.xw7yly9 > div > div.x9f619.x1n2onr6.x1ja2u2z.xeuugli.xs83m0k.xjl7jj.x1xmf6yo.x1emribx.x1e56ztr.x1i64zmx.x19h7ccj.xu9j1y6.x7ep2pv > div:nth-child(3)",
			CallBack: func(post act.FbPost) error {
				fmt.Println("Post message:", post.PostMessage)
				return nil
			},
		},
	}

	err := wba.RunWithCookiesFile(actions, fmt.Sprintf("%s_cookies.json", domain))
	require.NoError(t, err)
}

func TestQueryText(t *testing.T) {

	domain := "www.upwork.com"
	actions := []wba.IAction{
		&act.ActionGetFBPost{
			Url:       fmt.Sprintf("https://%s", domain),
			ClassPath: "#login > div > div > footer > div > div",
			CallBack: func(post act.FbPost) error {
				fmt.Println("Post message:", post.PostMessage)
				return nil
			},
		},
	}

	err := wba.RunWithCookiesFile(actions, fmt.Sprintf("%s_cookies.json", domain))
	require.NoError(t, err)
}
