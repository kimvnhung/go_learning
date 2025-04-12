package webbaseaction

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type ActionType int

const (
	ActNone      ActionType = iota // Do nothing
	ActOpenURL                     // Open a URL in the default browser
	ActGetFBPost                   // Get Facebook post
)

type IAction interface {
	GetActionType() ActionType     // Get the type of action
	Act(ctx context.Context) error // Perform the action
}

func cleanUp(cookieData []byte) []byte {
	log.Println("Cleaning up cookie data")
	var mapData []map[string]interface{}

	if err := json.Unmarshal(cookieData, &mapData); err != nil {
		log.Printf("Error unmarshalling JSON data: %v", err)
		return cookieData
	}

	// Loop in mapData in recursively, check if contains "sameSite" key and replace with below map
	// "none"-> "None"
	// "lax" -> "Lax"
	// "strict" -> "Strict"
	// "unspecified" -> "None"
	// "no_restriction" -> "None"
	for i := 0; i < len(mapData); i++ {
		if _, ok := mapData[i]["sameSite"]; ok {
			switch mapData[i]["sameSite"] {
			case "none":
				mapData[i]["sameSite"] = "None"
			case "lax":
				mapData[i]["sameSite"] = "Lax"
			case "strict":
				mapData[i]["sameSite"] = "Strict"
			case "unspecified":
				mapData[i]["sameSite"] = "None"
			case "no_restriction":
				mapData[i]["sameSite"] = "None"
			}
		}
	}

	// Marshal the modified data back to JSON
	cleanedData, err := json.Marshal(mapData)
	if err != nil {
		log.Printf("Error marshalling cleaned data: %v", err)
		return cookieData
	}
	log.Println("Cleaned up cookie data")
	return cleanedData
}

func LoadCookiesFromFile(filePath string) ([]*network.Cookie, error) {
	// Load file content
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	jsonData = cleanUp(jsonData)

	// Unmarshal JSON data into a slice of cookies
	var cookies []*network.Cookie
	if err := json.Unmarshal(jsonData, &cookies); err != nil {
		log.Printf("Error unmarshalling JSON data: %v", err)
		return nil, err
	}

	return cookies, nil
}

func Run(actions []IAction) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),                                // 👈 Run with GUI
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // Less bot-like
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"),
	)
	// Create a parent context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	log.Println("Prepared context")
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

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // 👈 Run with GUI
		// chromedp.Flag("disable-blink-features", "AutomationControlled"), // Less bot-like
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.6998.117 Safari/537.36"),
	)
	// Create a parent context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Load cookies into the context
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		for _, cookie := range cookies {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			err := network.SetCookie(cookie.Name, cookie.Value).
				WithDomain(cookie.Domain).
				WithPath(cookie.Path).
				WithExpires(&expr).
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
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),                                // 👈 Run with GUI
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // Less bot-like
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"),
	)
	// Create a parent context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
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
