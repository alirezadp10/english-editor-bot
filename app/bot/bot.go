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

    // Set commands for the bot
    commands := []telebot.Command{
        {Text: "check", Description: "با دستور /check میتونی متنت رو ویرایش کنی"},
        {Text: "formal", Description: "با دستور /formal میتونی شکل رسمی متنت رو ببینی"},
        {Text: "informal", Description: "با دستور /informal میتونی شکل خودمونی متنت رو ببینی"},
        {Text: "fa", Description: "با دستور /fa میتونی متنت رو به فارسی ترجمه کنی"},
        {Text: "en", Description: "با دستور /en میتونی متنت رو به انگلیسی ترجمه کنی"},
    }

    // Use the bot's setMyCommands method to apply these commands
    err = bot.SetCommands(commands)
    if err != nil {
        log.Fatal("Failed to set commands:", err)
    }

    log.Println("Commands successfully set for the bot!")

    return bot
}
