package main

// slack age bot

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.TimeStamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A043N1W504X-4138636833410-b7af698500e423d90ca47664a5fbd45c50a55516bcf8c4b0b94c8a30e468cade")
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4124055123623-4138664024402-IeuiBGlxt7c38kcuesmLwqq9")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2022 - yob
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
