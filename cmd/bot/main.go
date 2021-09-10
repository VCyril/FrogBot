package main

import (
	"frogBot/pkg/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	// This is test token, I know that keeping tokens in code is veeeery bad
	tgbot, err := tgbotapi.NewBotAPI("1967653234:AAE3p55A3x9qUrNXMRHAmRkNFEuJNYND3Vc")
	if err != nil {
		log.Fatalf("Error while creating BotApi: %s", err.Error())
	}
	tgbot.Debug = true

	client := telegram.NewClient()

	bot := telegram.NewBot(tgbot, client)
	err = bot.Start()
	if err != nil {
		log.Fatalf("Error while starting bot: %s", err.Error())
	}
}