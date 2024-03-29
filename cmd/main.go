package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/emilyzhang/lotto-alerts/pkg/scraper"
	"github.com/emilyzhang/lotto-alerts/pkg/sms"
)

// HandleRequest handles a lambda request.
func HandleRequest() ([]scraper.Lottery, error) {
	lotteries := scraper.Scrape("http://calottery.com")
	target, err := strconv.Atoi(os.Getenv("TARGET"))
	if err != nil {
		fmt.Println("Could not convert string to integer:", os.Getenv("TARGET"))
	}
	alert(lotteries, os.Getenv("PHONE_NUMBER"), target)
	return lotteries, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func alert(lotteries []scraper.Lottery, num string, target int) {
	weekday := int(time.Now().Weekday())
	var send bool
	for _, lottery := range lotteries {
		fmt.Println(lottery)
		if lottery.Millions > target {
			switch lottery.Name {
			case "Powerball", "Superlotto Plus":
				if weekday == 2 || weekday == 5 {
					send = true
				}
			case "Mega Millions":
				if weekday == 1 || weekday == 4 {
					send = true
				}
			}
			if send {
				sms.Send(lottery.String(), num)
			}
		}
		send = false
	}
}

// Old main function, before this was a lambda function.
// func main() {
// 	lotteries := scraper.Scrape("http://calottery.com")
// 	target, err := strconv.Atoi(os.Getenv("TARGET"))
// 	if err != nil {
// 		fmt.Println("Could not convert string to integer:", os.Getenv("TARGET"))
// 	}
// 	alert(lotteries, os.Getenv("PHONE_NUMBER"), target)
// }
