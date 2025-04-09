package webbaseaction

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
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

func LoadCookiesFromFile(filePath string) ([]*network.Cookie, error) {
	// Load file content
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a slice of cookies
	var cookies []*network.Cookie
	if err := json.Unmarshal(jsonData, &cookies); err != nil {
		return nil, err
	}

	return cookies, nil
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

func RunWithCookiesFile(actions []IAction, filePath string) error {
	// Load cookies from file
	cookies, err := LoadCookiesFromFile(filePath)
	if err != nil {
		return err
	}

	// Create a parent context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Load cookies into the context
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		for _, cookie := range cookies {
			expireTime := &cdp.TimeSinceEpoch{}
			expireTime.UnmarshalJSON([]byte(fmt.Sprintf("%d", int(cookie.Expires))))
			err := network.SetCookie(cookie.Name, cookie.Value).
				WithDomain(cookie.Domain).
				WithPath(cookie.Path).
				WithExpires(expireTime).
				Do(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})); err != nil {
		return err
	}

	for _, action := range actions {
		if err := action.Act(ctx); err != nil {
			return err
		}
	}
	return nil
}

func RunWithCookies(actions []IAction, cookies map[string]string) error {
	// Create a parent context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Load cookies into the context
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		for key, value := range cookies {
			err := network.SetCookie(key, value).
				Do(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})); err != nil {
		return err
	}

	for _, action := range actions {
		if err := action.Act(ctx); err != nil {
			return err
		}
	}
	return nil
}
