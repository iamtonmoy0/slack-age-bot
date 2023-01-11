package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {
	os.Setenv("slack bot token", "xoxb-4641079431152-4620665186933-Kxo2K1y7hcfwxWcHLIm0XXeq")
	os.Setenv("slack app token", "xapp-1-A04JQ6MQYQ1-4636230567041-00abef55db3ef05f1dacded3c8cc82c9899c3320e6eba652241355bec0434cd3")

	bot := slacker.NewClient(os.Getenv("slack bot token"), os.Getenv("slack app token"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob  is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",

		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
