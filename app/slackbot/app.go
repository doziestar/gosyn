package slackbot

import (
	"context"
	"fmt"
	"os"

	"github.com/shomali11/slacker"
)

func Init() {
	// Start the bot up
	// Start()
	fmt.Println("Hello, Slack bot.")
	os.Setenv("SLACK_API_TOKEN", "xapp-1-A03F0RTJ9RT-3507154844982-c2e64558013cdde6d75f6951f93db9eb952c056e85a59d9cc17b6324d0a096c2")
	os.Setenv("SLACK_BOT_TOKEN", "xapp-1-A03F0RTJ9RT-3507154844982-c2e64558013cdde6d75f6951f93db9eb952c056e85a59d9cc17b6324d0a096c2")

	api := slacker.NewClient(os.Getenv("SLACK_API_TOKEN"), os.Getenv("SLACK_BOT_TOKEN"))
	// api.Command("/hello", helloHandler)
	fmt.Println(api.BotCommands())

	go printCommands(api.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := api.Listen(ctx)
	if err != nil {
		fmt.Println(err)
	}

	api.Command("/help", &slacker.CommandDefinition{
		Description: "This is a help command",
		Example:     "/help",
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Welcome! I am a bot. How may i help you?")
		},
	})
	api.Command("/form", &slacker.CommandDefinition{
		Description: "Form",
		Example:     "/form",
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Form")
		},
	})

}

// func helloHandler(request *slacker.Request, response *slacker.ResponseWriter) {
// 	// response.Reply("Hello, world!")
// }

func printCommands(analytics <-chan *slacker.CommandEvent) {
	for event := range analytics {
		fmt.Println("Command:", event.Command)
		fmt.Println(event.Command)
		fmt.Println(event.Timestamp)
	}
}
