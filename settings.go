package main

import (
	"os"
	"strconv"
)

var PORT = os.Getenv("PORT")
var BOT_TOKEN = os.Getenv("BOT_TOKEN")
var ADMIN_CHAT_ID, _ = strconv.ParseInt(os.Getenv("ADMIN_CHAT_ID"), 10, 64)
