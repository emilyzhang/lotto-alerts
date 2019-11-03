package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Lottery represents a lottery.
type Lottery struct {
	Name      string
	DrawDate  string
	Value     string
	CashValue string
	Millions  int
}

// String returns an html string of lottery information.
func (l *Lottery) String() string {
	return l.Name + "<br>" + l.DrawDate + "<br>" + "Value: " + l.Value + "<br>" + l.CashValue
}

// Scrape scrapes the CA lottery website for information about current lotteries.
func Scrape(website string) []Lottery {
	var err error
	lotteries := make([]Lottery, 0)
	c := colly.NewCollector()

	c.OnHTML("#draw-game-listing", func(e *colly.HTMLElement) {
		e.ForEach(".card", func(_ int, el *colly.HTMLElement) {
			name := strings.Split(el.ChildText(".card-header"), " ")
			g := Lottery{
				Name:      strings.Join(name[:len(name)-2], " "),
				CashValue: el.ChildText(".draw-cards--cash-value"),
				DrawDate:  el.ChildText(".draw-cards--next-draw-date"),
			}

			lotteryValue := strings.Split(el.ChildText(".draw-cards--lottery-amount"), " ")
			g.Value = strings.Join(lotteryValue, " ")
			if lotteryValue[len(lotteryValue)-1] == "MILLION*" {
				g.Millions, err = strconv.Atoi(lotteryValue[0][1:])
				if err != nil {
					fmt.Println("issues converting from string to int")
				}
			}
			lotteries = append(lotteries, g)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(website)
	return lotteries
}
