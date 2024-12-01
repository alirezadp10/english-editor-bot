package main

import (
    "english-editor-bot/bot"
    "english-editor-bot/database"
    "fmt"
    "log"
    "os"
)

func main() {
    // Connect to the database
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Database connection error: %v", err)
    }

    _ = database.AutoMigrate(db)

    // Initialize bot and start handlers
    botToken := os.Getenv("BOT_TOKEN")
    apiKey := os.Getenv("API_KEY")

    telegramBot := bot.InitializeBot(botToken)
    bot.SetupHandlers(telegramBot, db, apiKey)

    fmt.Println("Bot is running...")
    telegramBot.Start()
}
