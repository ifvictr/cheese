package cheese

import (
	"fmt"

	"github.com/slack-go/slack"
)

func OnMessage(slackClient *slack.Client, event *slack.MessageEvent) {
	// User must not be a bot
	isValidUser := event.BotID == "" && event.User != "USLACKBOT" && event.User != ""
	// Must be a new message on not another subtype (e.g. message_changed)
	isNewMessage := event.SubType == "" || event.SubType == "message_replied"
	if !isValidUser || !isNewMessage || !HasCheeseTouch(event.User) || !HasIntentToPass(event.Text) {
		return
	}

	// The user has the cheese touch and intends to pass it.
	res, err := GetPrecedingMessage(slackClient, event.Channel, event.TimeStamp)
	if err != nil || len(res.Messages) == 0 {
		fmt.Println("Failed to get the preceding message")
		slackClient.PostEphemeral(event.Channel, event.User, slack.MsgOptionText(
			PassErrorMessage, false))
		return
	}

	// Make sure the users are different.
	precedingUserId := res.Messages[0].User
	if precedingUserId == event.Username {
		return
	}

	// Check if the other user is safe.
	if HasFingersCrossed(res.Messages[0].Text) {
		slackClient.PostEphemeral(event.Channel, event.User, slack.MsgOptionText(
			fmt.Sprintf(FingersCrossedMessage, precedingUserId), false))
		return
	}

	// Pass the cheese touch to the other user.
	GiveCheeseTouch(precedingUserId)
	slackClient.PostMessage(event.Channel, slack.MsgOptionText(
		fmt.Sprintf(PassSuccessMessage, event.User, precedingUserId), false))
}

func OnReactionAdded(slackClient *slack.Client, event *slack.ReactionAddedEvent) {
	// Listen for users reacting with a pointing_up emoji.
	if event.Reaction != "point_up" && event.Reaction != "point_up_2" {
		return
	}

	// If there's already a cheese touch, don't look for users who "touch" a message
	// with a cheese emoji in it via reaction.
	if HasCheeseTouchStarted() {
		return
	}

	// Check for the cheese emoji.
	res, err := GetMessage(slackClient, event.Item.Channel, event.Item.Timestamp)
	if err != nil || len(res.Messages) == 0 {
		fmt.Println("Failed to get the message of reaction")
		return
	}
	if !HasCheese(res.Messages[0].Text) {
		return
	}

	// There's no cheese touch yet, so this is the first interaction a user has
	// had with the cheese. This user gets the cheese touch!
	GiveCheeseTouch(event.User)
	slackClient.PostMessage(event.Item.Channel, slack.MsgOptionText(
		fmt.Sprintf(StartedMessage, event.User), false))
}
