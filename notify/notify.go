package notify

import (
	"gopkg.in/telegram-bot-api.v4"
)

type Notifier interface {
	Send(s string) error
}

var _ Notifier = new(TelegramNotifier)

type TelegramNotifier struct {
	bot    *tgbotapi.BotAPI
	token  string
	chatID int64
}

func NewTelegramNotifier(token string, chatID int64) (*TelegramNotifier, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramNotifier{
		bot:    bot,
		token:  token,
		chatID: chatID,
	}, nil
}

func (t *TelegramNotifier) Send(s string) error {
	msg := tgbotapi.NewMessage(t.chatID, s)
	_, err := t.bot.Send(msg)
	return err
}
