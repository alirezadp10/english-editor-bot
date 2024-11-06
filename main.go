package main

import (
    "english-editor-bot/bot"
    "english-editor-bot/config"
    "english-editor-bot/database"
    "fmt"
    "log"
    "os"
)

func main() {
    // Set up a new Kafka producer
    //config := sarama.NewConfig()
    //config.Producer.RequiredAcks = sarama.WaitForAll
    //config.Producer.Retry.Max = 5
    //config.Producer.Return.Successes = true
    ////
    //producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
    //if err != nil {
    //    panic(err)
    //}
    //defer producer.Close()
    //
    //// Create a message
    //msg := &sarama.ProducerMessage{
    //    Topic: "messages",
    //    Value: sarama.StringEncoder(`{"user_id":123,"content":"Hello, world!","timestamp":"2024-11-05T12:34:56Z"}`),
    //}
    //
    //// Send message
    //partition, offset, err := producer.SendMessage(msg)
    //if err != nil {
    //    panic(err)
    //}
    //fmt.Printf("Message sent to partition %d with offset %d\n", partition, offset)

    //fmt.Printf("Message sent to partition with offset\n")

    // Load environment variables
    config.LoadEnv()

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
