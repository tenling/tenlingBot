package main

import (
	"log"

	"github.com/tenling/tenlingBot/Action"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("test")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msgText := ""
		switch update.Message.Text {
		case "/bank", "/conanBank", "/帳本":
			msgText = Action.ReplyBank()
		case "/女朋友", "/Henry's girlfriend":
			msgText = "nil"
		default:
			msgText = update.Message.Text
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
