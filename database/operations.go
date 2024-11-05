package database

import (
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
)

func SaveUser(db *gorm.DB, user *telebot.User) (*User, error) {
    newUser := User{
        ID:           user.ID,
        FirstName:    user.FirstName,
        LastName:     user.LastName,
        Username:     user.Username,
        LanguageCode: user.LanguageCode,
        IsBot:        user.IsBot,
        IsPremium:    user.IsPremium,
    }
    err := db.Save(&newUser).Error
    return &newUser, err
}

func SaveChat(db *gorm.DB, chat *telebot.Chat) (*Chat, error) {
    newChat := Chat{
        ID:                      chat.ID,
        Type:                    chat.Type,
        Title:                   chat.Title,
        FirstName:               chat.FirstName,
        LastName:                chat.LastName,
        Username:                chat.Username,
        CanSendPaidMedia:        chat.CanSendPaidMedia,
        HasVisibleHistory:       chat.HasVisibleHistory,
        BackgroundCustomEmojiID: chat.BackgroundCustomEmojiID,
    }
    err := db.Save(&newChat).Error
    return &newChat, err
}

func SaveMessage(db *gorm.DB, message *telebot.Message, user *User, chat *Chat, repliedMessage *ReplyToMessage) (*Message, error) {
    newMessage := Message{
        MessageID:        message.ID,
        MessageThreadID:  message.ThreadID,
        FromID:           user.ID,
        Date:             message.Unixtime,
        ChatID:           chat.ID,
        ReplyToMessageID: repliedMessage.MessageID,
        Text:             message.Text,
    }
    err := db.Save(&newMessage).Error
    return &newMessage, err
}

func SaveReplyToMessage(db *gorm.DB, repliedMessage *telebot.Message, user *User, chat *Chat) (*ReplyToMessage, error) {
    if repliedMessage == nil {
        return nil, nil
    }

    newReply := ReplyToMessage{
        MessageID:       repliedMessage.ID,
        MessageThreadID: repliedMessage.ThreadID,
        UserID:          user.ID,
        Date:            repliedMessage.Unixtime,
        Text:            repliedMessage.Text,
        ChatID:          chat.ID,
    }
    err := db.Save(&newReply).Error
    return &newReply, err
}

func SaveEntities(db *gorm.DB, entities []telebot.MessageEntity, messageID int) error {
    for _, entity := range entities {
        newEntity := Entity{
            Type:      entity.Type,
            Offset:    int64(entity.Offset),
            Length:    int64(entity.Length),
            MessageID: messageID,
        }
        if err := db.Save(&newEntity).Error; err != nil {
            return err
        }
    }
    return nil
}
