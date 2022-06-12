package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

type Bot struct {
	AdminChatID int64
	Api         *tgbotapi.BotAPI
}

func (bot *Bot) Init(token string) error {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	bot.Api = api

	bot.Api.Send(
		tgbotapi.NewSetMyCommands(
			tgbotapi.BotCommand{
				"start",
				"Стартует бота",
			},
			tgbotapi.BotCommand{
				"help",
				"Информация о боте",
			},
		),
	)
	return nil
}

func (bot *Bot) setAdmin(chatID int64) {
	bot.AdminChatID = chatID
}

func (bot *Bot) forwardToAdmin(message tgbotapi.Message) {
	forward := tgbotapi.NewCopyMessage(bot.AdminChatID, message.Chat.ID, message.MessageID)
	bot.Api.Send(forward)

	reply_info := tgbotapi.NewMessage(bot.AdminChatID, "От @"+message.Chat.UserName)
	bot.Api.Send(reply_info)

	reply := tgbotapi.NewMessage(message.Chat.ID, "Cпасибо, отправил админу 😉")
	bot.Api.Send(reply)
}

func (bot *Bot) handleCommand(command string, chatID int64) {
	switch command {
	case "start", "help":
		if chatID == bot.AdminChatID {
			message := tgbotapi.NewMessage(chatID, StartMessageAdminTemplate())
			bot.Api.Send(message)
		} else {
			message := tgbotapi.NewMessage(chatID, StartMessageTemplate())
			bot.Api.Send(message)
		}
	default:
		message := tgbotapi.NewMessage(chatID, "Неизвечтная команда. Используй /help чтобы узнать известные мне команды.")
		bot.Api.Send(message)
	}
}

func (bot *Bot) Run() error {
	if bot.Api == nil {
		return os.ErrNotExist
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.Api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch {
		case update.Message.IsCommand():
			bot.handleCommand(update.Message.Command(), update.Message.Chat.ID)
		default:
			bot.forwardToAdmin(*update.Message)
		}

	}
	return nil
}
