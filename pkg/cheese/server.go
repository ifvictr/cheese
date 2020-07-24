package cheese

import (
	"github.com/go-redis/redis"
	"github.com/slack-go/slack"
)

var redisClient *redis.Client

func StartServer(config *Config) {
	// Set up Redis connection
	options, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(options)

	// Initialize Slack app
	api := slack.New(config.BotToken)
	rtm := api.NewRTM()

	go rtm.ManageConnection()

	// Listen to and forward events to their handlers, if they exist.
	for event := range rtm.IncomingEvents {
		switch e := event.Data.(type) {
		case *slack.MessageEvent:
			OnMessage(&rtm.Client, e)
			continue
		case *slack.ReactionAddedEvent:
			OnReactionAdded(&rtm.Client, e)
			continue
		}
	}
}
