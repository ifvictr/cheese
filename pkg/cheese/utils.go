package cheese

import (
	"strings"

	"github.com/slack-go/slack"
)

const (
	CheeseEmoji = ":cheese_wedge:"
	PassEmoji1  = ":point_up:"
	PassEmoji2  = ":point_up_2:"
	SafeEmoji   = ":crossed_fingers:"
)

func GetMessage(slackClient *slack.Client, channelId string, messageId string) (*slack.GetConversationHistoryResponse, error) {
	return slackClient.GetConversationHistory(&slack.GetConversationHistoryParameters{
		ChannelID: channelId,
		Inclusive: true,
		Latest:    messageId,
		Limit:     1,
	})
}

func GetPrecedingMessage(slackClient *slack.Client, channelId string, fromMessageId string) (*slack.GetConversationHistoryResponse, error) {
	return slackClient.GetConversationHistory(&slack.GetConversationHistoryParameters{
		ChannelID: channelId,
		Latest:    fromMessageId,
		Limit:     1,
	})
}

func HasIntentToPass(message string) bool {
	return strings.Contains(message, PassEmoji1) || strings.Contains(message, PassEmoji2)
}

func HasFingersCrossed(message string) bool {
	return strings.Contains(message, SafeEmoji)
}

func HasCheese(message string) bool {
	return strings.Contains(message, CheeseEmoji)
}

func HasCheeseTouch(userId string) bool {
	bearingUserId, err := redisClient.Get("bearing_user_id").Result()

	if err != nil {
		return false
	}

	return bearingUserId == userId
}

func HasCheeseTouchStarted() bool {
	exists, err := redisClient.Exists("bearing_user_id").Result()

	if err != nil {
		return false
	}

	return exists == 1
}

func GiveCheeseTouch(userId string) {
	redisClient.Set("bearing_user_id", userId, 0)
}

func IsPublicChannel(channelId string) bool {
	return strings.HasPrefix(channelId, "C")
}
