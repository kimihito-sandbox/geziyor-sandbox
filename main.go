package main

import (
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			req, _ := client.NewRequest("GET", "https://www.e-uchina.net/jukyo?bukken_type=3&sort=senyu_space_large", nil)
			req.Rendered = true
			req.Actions = []chromedp.Action{
				chromedp.Navigate("https://www.e-uchina.net/jukyo?bukken_type=3&sort=senyu_space_large"),
				chromedp.WaitReady("div.search-result"),
				chromedp.ActionFunc(func(ctx context.Context) error {
					node, err := dom.GetDocument().Do(ctx)
					if err != nil {
						return err
					}
					_, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
					return err
				}),
			}
			g.Do(req, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.search-result-item").Each(func(_ int, s *goquery.Selection) {
				fmt.Println(s.Find("div.result-content").Text())
			})
		},
	}).Start()
}
