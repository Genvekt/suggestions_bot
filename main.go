package main

import (
	"fmt"
	"log"
	"net/http"
	"suggestion_bot/bot"
)

func main() {
	SuggestionBot := bot.Bot{
		ADMIN_CHAT_ID,
		nil,
	}
	fmt.Println("Initializing bot...")
	if err := SuggestionBot.Init(BOT_TOKEN); err != nil {
		panic(err)
	}
	go SuggestionBot.Run()
	fmt.Println("Bot started.")
	fmt.Println("Listening for http...")

	http.HandleFunc("/", handler) // each request calls handler
	fmt.Println("")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
