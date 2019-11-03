package main

import (
	"fmt"
	"os"

	"github.com/emilyzhang/lotto-alerts/pkg/scraper"
	"github.com/emilyzhang/lotto-alerts/pkg/sms"
)

func main() {
	games := scraper.Scrape("http://calottery.com")
	fmt.Println(games)
	sms.Send(games[0].String(), os.Getenv("PHONE_NUMBER"))
}
