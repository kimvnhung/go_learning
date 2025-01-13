package crawner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
)

type Cookier struct {
	CookiePath string
	ctx        context.Context
	cancel     context.CancelFunc
	tasks      chromedp.Tasks
}

const (
	indexHTML = `<!doctype html>
<html>
<body>
  <div id="result">%s</div>
</body>
</html>`
)

// setcookies returns a task to navigate to a host with the passed cookies set
// on the network request.
func setcookies(host string, res *string, cookies ...string) chromedp.Tasks {
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for i := 0; i < len(cookies); i += 2 {
				err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain("localhost").
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		// navigate to site
		chromedp.Navigate(host),
		// read the returned values
		chromedp.Text(`#result`, res, chromedp.ByID, chromedp.NodeVisible),
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := storage.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}

func cookieServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookies := req.Cookies()
		for i, cookie := range cookies {
			log.Printf("from %s, server received cookie %d: %v", req.RemoteAddr, i, cookie)
		}
		buf, err := json.MarshalIndent(req.Cookies(), "", "  ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(res, indexHTML, string(buf))
	})
	return http.ListenAndServe(addr, mux)
}

func New() *Cookier {
	go cookieServer("localhost:9222")

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())

	cookier := &Cookier{
		CookiePath: "cookie.json",
		ctx:        ctx,
		cancel:     cancel,
	}

	return cookier
}

func (c *Cookier) Run() {
	// Load the cookies from the file then put to

	defer c.cancel()
	if err := chromedp.Run(c.ctx, c.tasks); err != nil {
		log.Fatal(err)
	}
}
