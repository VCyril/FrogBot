package telegram


import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"time"
)

const (
	updateTimeout = 60
)


type Bot struct {
	bot *tgbotapi.BotAPI
	client *HTTPClient
}

type HTTPClient struct {
	client *http.Client

}

func NewClient() *HTTPClient{
	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func NewBot(bot *tgbotapi.BotAPI, client *HTTPClient) *Bot {
	return &Bot{
		bot: bot,
		client: client,
	}
}


func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.getUpdatesChan()
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) getUpdatesChan() (tgbotapi.UpdatesChannel, error){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = updateTimeout

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	return updates, nil
}