package bot

import (
    "gopkg.in/telebot.v4"
    "log"
    "time"
)

func InitializeBot(botToken string) *telebot.Bot {
    pref := telebot.Settings{
        Token:  botToken,
        Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
    }
    bot, err := telebot.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }
    return bot
}
