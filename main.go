package main

import (
	"log"
	"strings"

	"gopkg.in/telegram-bot-api.v4"

	"github.com/ikurakin/translate-bot/config"
	"github.com/ikurakin/translate-bot/translate"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	translator := translate.New(cfg.TranslateApiKey, cfg.TranslateApiURL)
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotKey)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		result, err := translator.Translate(cfg.LangugeSrc, cfg.LanguageDst, update.Message.Text)
		if err != nil {
			panic(err)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(result, "\n"))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
