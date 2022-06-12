package main

import "suggestion_bot/bot"

func main() {
	SuggestionBot := bot.Bot{
		ADMIN_CHAT_ID,
		nil,
	}

	if err := SuggestionBot.Init(BOT_TOKEN); err != nil {
		panic(err)
	}

	if err := SuggestionBot.Run(); err != nil {
		panic(err)
	}
}
