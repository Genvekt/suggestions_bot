package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"suggestion_bot/bot"
)

func main() {
	SuggestionBot := bot.Bot{
		ADMIN_CHAT_ID,
		nil,
	}
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler) // each request calls handler

	if err := SuggestionBot.Init(BOT_TOKEN); err != nil {
		panic(err)
	}

	if err := SuggestionBot.Run(); err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
