package main

import (
	"fmt"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analysisChannel <-chan *slacker.CommandEvent) {
	for event := range analysisChannel {
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
}
