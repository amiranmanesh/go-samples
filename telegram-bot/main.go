package main

import (
	//"log"
	//"time"
	//
	//"github.com/go-telegram-bot-api/telegram-bot-api"

	//"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	_ = godotenv.Load()

	poller := &tb.LongPoller{Timeout: 10 * time.Second}
	spamProtected := tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
		if upd.Message == nil {
			return true
		}

		if strings.Contains(upd.Message.Text, "spam") {
			return false
		}

		return true
	})

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "",

		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: spamProtected,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		// Universal markup builders.
		menu     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		selector = &tb.ReplyMarkup{}

		// Reply buttons.
		btnHelp     = menu.Text("ℹ Help")
		btnSettings = menu.Text("⚙ Settings")

		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind,
		// as it has to be for callback routing to work.
		//
		btnPrev = selector.Data("⬅", "prev", "data")
		btnNext = selector.Data("➡", "next", "data")
	)

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	// Command: /start <PAYLOAD>
	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		b.Send(m.Sender, "Hello!", menu)
	})

	// On reply button pressed (message)
	b.Handle(&btnHelp, func(m *tb.Message) {
		b.Send(m.Sender, "help", selector)
	})

	// On inline button pressed (callback)
	b.Handle(&btnPrev, func(c *tb.Callback) {
		// ...
		// Always respond!

		b.Respond(c, &tb.CallbackResponse{
			Text:      "temp",
			ShowAlert: false,
		})
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers

		b.Send(m.Sender, "Hello 1")
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		// photos only
		b.Send(m.Sender, "Hello 2")
	})

	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		// channel posts only
		b.Send(m.Sender, "Hello 3")
	})

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		// incoming inline queries

	})

	//b.Handle("/start", func(m *tb.Message) {
	//	if !m.Private() {
	//		return
	//	}
	//
	//	fmt.Println(m.Payload) // <PAYLOAD>
	//	b.Send(m.Sender, "Hello 4")
	//})
	b.Start()
	//bot, err := tgbotapi.NewBotAPI("1619145503:AAFKpUbGqVtkIGXzjNc00sl_JaqTGuqsHRc")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//bot.Debug = true
	//
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 60
	//
	//updates, err := bot.GetUpdatesChan(u)
	//
	//for update := range updates {
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//	}
	//
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	msg.ReplyToMessageID = update.Message.MessageID
	//
	//	bot.Send(msg)
	//}
}
