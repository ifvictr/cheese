package cheese

// Raw format strings for messages to use with sprintf.
const (
	FingersCrossedMessage string = "Looks like <@%s> had their fingers crossed. Too bad!"
	PassErrorMessage      string = "I couldn’t check if the Cheese Touch could be passed to the user above you. Please try again later."
	PassSuccessMessage    string = CheeseEmoji + " <@%s> passed the Cheese touch to <@%s>!"
	StartedMessage        string = ":eyes: :rotating_light: *<@%s>* just touched some terribly moldy cheese, starting the Cheese Touch! To pass it on to someone else, reply with either " +
		PassEmoji1 + " or a " + PassEmoji2 + " under someone who doesn’t have their fingers crossed (a " +
		SafeEmoji + " in their message)."
)
