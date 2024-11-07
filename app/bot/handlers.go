package bot

import (
    "english-editor-bot/api"
    "english-editor-bot/database"
    "fmt"
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
)

func SetupHandlers(bot *telebot.Bot, db *gorm.DB, apiKey string) {
    bot.Handle("/check", func(c telebot.Context) error {
        return handleTextMessage(c, db, apiKey)
    })
}

func handleTextMessage(c telebot.Context, db *gorm.DB, apiKey string) error {
    // Database operations for saving user, chat, message, etc.
    user, err := database.SaveUser(db, c.Message().Sender)
    if err != nil {
        return fmt.Errorf("failed to save user: %v", err)
    }

    chat, err := database.SaveChat(db, c.Message().Chat)
    if err != nil {
        return fmt.Errorf("failed to save chat: %v", err)
    }

    var repliedMessage *database.ReplyToMessage
    if c.Message().ReplyTo != nil {
        repliedUser, err := database.SaveUser(db, c.Message().ReplyTo.Sender)
        if err != nil {
            return fmt.Errorf("failed to save replied user: %v", err)
        }

        repliedMessage, err = database.SaveReplyToMessage(db, c.Message().ReplyTo, repliedUser, chat)
        if err != nil {
            return fmt.Errorf("failed to save replied message: %v", err)
        }

        message, err := database.SaveMessage(db, c.Message(), user, chat, repliedMessage)
        if err != nil {
            return fmt.Errorf("failed to save message: %v", err)
        }

        if err := database.SaveEntities(db, c.Message().Entities, message.MessageID); err != nil {
            return fmt.Errorf("failed to save entities: %v", err)
        }
    }

    if repliedMessage == nil {
        return c.Reply("باید رو مسیجی که می‌خوای تصحیح شه ریپلای کنی", telebot.ModeHTML)
    }

    // Request creation and API call
    requestBody := api.CreateRequestBody(repliedMessage.Text)
    responseBody, err := api.SendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    // Parse and reply
    content, err := api.ParseResponse(responseBody)
    if err != nil {
        return err
    }
    return c.Reply(content, telebot.ModeHTML)
}
