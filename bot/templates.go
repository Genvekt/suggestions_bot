package bot

func StartMessageTemplate() string {
	return "Я бот предложки смешных постов. \nЕсли у тебя есть смешные картинки/видео, отправляй мне. \n\nИзвестные команды:\n/start /help - информация обо мне."
}

func StartMessageAdminTemplate() string {
	return "Я бот предложки смешных постов. \nА ты мой админ. Жди предложек! \n\nИзвестные команды:\n/start /help - информация обо мне."
}
