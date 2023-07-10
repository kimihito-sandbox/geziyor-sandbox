package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.e-uchina.net/jukyo?bukken_type=3&sort=senyu_space_large", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.search-result-item").Each(func(_ int, s *goquery.Selection) {
				fmt.Println(s.Find("div.result-content").Text())
			})
		},
	}).Start()
}
