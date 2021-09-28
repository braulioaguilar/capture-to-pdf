package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var (
	urlFlag  *string
	nameFlag *string
)

func main() {
	urlFlag = flag.String("url", "https://www.google.com/", "URL to get the info")
	nameFlag = flag.String("name", "example.pdf", "The filename of your file")

	flag.Parse()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(*urlFlag, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(*nameFlag, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// printToPDF -- Prints a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
