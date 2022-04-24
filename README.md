# SeagalBot in Go

This is a re-write of [the original SeagalBot](https://github.com/ItsBradyDavis/SeagalBot), just in Go. All creativity attributed to the original author.

## Local development

1. [Install Go](https://go.dev/doc/install)
2. `go mod download` - install dependencies
3. Copy `.env.example` as `.env` and add `BOT_TOKEN` and `GIPHY_API_KEY`
   1. See [this tutorial](https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf) for creating a bot account and the `BOT_TOKEN`. Also follow these instructions for linking the bot to your server.
   2. See [this article](https://support.giphy.com/hc/en-us/articles/360020283431-Request-A-GIPHY-API-Key) for requesting a Giphy API key
4. Authorize the bot
5. `go build` - compile the application
6. `./seagal-bot-go`
