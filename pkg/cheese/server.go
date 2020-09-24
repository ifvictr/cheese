package cheese

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type slashCommandResponse struct {
	Blocks       slack.Blocks `json:"blocks,omitempty"`
	Text         string       `json:"text,omitempty"`
	ResponseType string       `json:"response_type,omitempty"`
}

func (resp slashCommandResponse) MarshalJSON() []byte {
	marshalled, _ := json.Marshal(resp)
	return marshalled
}

var (
	cheeseConfig *Config
	redisClient  *redis.Client
	slackClient  *slack.Client
)

func StartServer(config *Config) {
	cheeseConfig = config
	// Set up Redis connection
	options, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(options)

	// Initialize Slack app
	slackClient = slack.New(config.BotToken)

	// Start receiving events
	http.HandleFunc("/slack/events", handleSlackEvents)
	http.HandleFunc("/slack/cheesewho", handleSlashCommand)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}

func handleSlackEvents(w http.ResponseWriter, r *http.Request) {
	// Verify the payload was sent by Slack.
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	apiEvent, err := slackevents.ParseEvent(json.RawMessage(body),
		slackevents.OptionVerifyToken(
			&slackevents.TokenComparator{VerificationToken: cheeseConfig.VerificationToken}))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Handle the event that came through
	switch apiEvent.Type {
	case slackevents.URLVerification:
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
		break
	case slackevents.CallbackEvent:
		HandleInnerEvent(slackClient, &apiEvent.InnerEvent)
		break
	}
}

func handleSlashCommand(w http.ResponseWriter, r *http.Request) {
	command, _ := slack.SlashCommandParse(r)

	if !command.ValidateToken(cheeseConfig.VerificationToken) {
		w.WriteHeader(400)
		w.Write([]byte("Request not verified."))
		return
	}

	bearerOfCheese := WhoHasCheeseTouch()

	var text string

	if bearerOfCheese == "" {
		text = "Looks like nobody has the Cheese Touch... yet... :cheese_wedge:"
	} else if bearerOfCheese == command.UserID {
		text = "Oh no, _you_ have the Cheese Touch! :cheese_wedge: Give it to someone else by reacting to their message with :point_up:!"
	} else {
		text = fmt.Sprintf("Looks like <@%s> has the Cheese Touch right now! :cheese_wedge:", bearerOfCheese)
	}

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", text, false, false), nil, nil),
		},
	}

	resp := slashCommandResponse{
		ResponseType: "ephemeral",
		Blocks:       blocks,
	}

	marshalled := resp.MarshalJSON()

	w.Header().Add("Content-type", "application/json")
	w.Write(marshalled)
}
