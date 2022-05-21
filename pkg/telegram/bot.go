package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	AddProductCommand    = ""
	RemoveProductCommand = "remove"
	ListProductsCommand  = "list"
	ResetCartCommand     = "reset"
)

type messageProcessor func(int64, string) string

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Run() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.processMessage(update)
	}

	return nil
}

func (b *Bot) processMessage(update tgbotapi.Update) {
	f, exists := processors[update.Message.Command()]

	if !exists {
		return
	}

	response := f(update.Message.Chat.ID, update.Message.Text)

	if response == "" {
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}
