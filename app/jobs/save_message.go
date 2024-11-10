package jobs

import (
    "english-editor-bot/database"
    "fmt"
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
)

func SaveMessage(c telebot.Context, db *gorm.DB) error {
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
    return nil
}
