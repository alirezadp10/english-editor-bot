package database

import (
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
)

type User struct {
    ID           int64  `gorm:"primaryKey"`
    FirstName    string `gorm:"column:first_name"`
    LastName     string `gorm:"column:last_name"`
    Username     string
    LanguageCode string `gorm:"column:language_code"`
    IsBot        bool   `gorm:"column:is_bot"`
    IsPremium    bool   `gorm:"column:is_premium"`
}

type Chat struct {
    ID                      int64 `gorm:"primaryKey"`
    Type                    telebot.ChatType
    Title                   string
    FirstName               string `gorm:"column:first_name"`
    LastName                string `gorm:"column:last_name"`
    Username                string
    CanSendPaidMedia        bool   `gorm:"column:can_send_paid_media"`
    HasVisibleHistory       bool   `gorm:"column:has_visible_history"`
    BackgroundCustomEmojiID string `gorm:"column:background_custom_emoji_id"`
}

type ReplyToMessage struct {
    MessageID       int `gorm:"primaryKey"`
    MessageThreadID int
    UserID          int64
    Date            int64
    Text            string
    User            User  `gorm:"foreignKey:UserID"`
    Chat            Chat  `gorm:"foreignKey:ChatID"`
    ChatID          int64 `gorm:"column:chat_id"`
}

type Message struct {
    MessageID        int `gorm:"primaryKey"`
    MessageThreadID  int
    FromID           int64 `gorm:"column:from_id"`    // ID for the user who sent the message
    From             User  `gorm:"foreignKey:FromID"` // Specify foreign key as FromID
    Date             int64
    ChatID           int64           `gorm:"column:chat_id"`
    Chat             Chat            `gorm:"foreignKey:ChatID"`
    ReplyToMessageID int             `gorm:"column:reply_to_message_id"`
    ReplyToMessage   *ReplyToMessage `gorm:"foreignKey:ReplyToMessageID"`
    Text             string
    Entities         []Entity `gorm:"foreignKey:MessageID"`
}

type Entity struct {
    ID        int64 `gorm:"primaryKey"`
    Type      telebot.EntityType
    Offset    int64
    Length    int64
    MessageID int `gorm:"column:message_id"`
}

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(&User{}, &Chat{}, &ReplyToMessage{}, &Message{}, &Entity{})
}
