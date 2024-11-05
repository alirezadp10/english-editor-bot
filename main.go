package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/joho/godotenv"
    "gopkg.in/telebot.v4"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

type Response struct {
    Choices []Choice `json:"choices"`
}

type Choice struct {
    Message Message `json:"message"`
}

type Message struct {
    Content string `json:"content"`
}

const (
    apiURL     = "https://api.deepinfra.com/v1/openai/chat/completions"
    systemRole = "Act as an English teacher to check and correct the sentences in your responses. Provide the revised version, list corrected errors in bullet form, and explain the corrections in English. For example, if the input is: \"Does it make sense being sad for being left by a partner who has threatened you?\" your answer format should be as follows, without additional introductory text or section titles:\n\n<b>✅ [Corrected form]</b>\n\n▶️ [Meaning or explanation]\n\n<blockquote>[Corrected errors in bullet list]</blockquote>"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    botToken := os.Getenv("BOT_TOKEN") // Get bot token from environment variables
    apiKey := os.Getenv("API_KEY")     // Get API key from environment variables

    bot := initializeBot(botToken)
    setupHandlers(bot, apiKey)

    fmt.Println("Bot is running...")
    bot.Start()
}

func initializeBot(botToken string) *telebot.Bot {
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

func setupHandlers(bot *telebot.Bot, apiKey string) {
    bot.Handle("/check", func(c telebot.Context) error {
        return handleTextMessage(c, apiKey)
    })
}

func handleTextMessage(c telebot.Context, apiKey string) error {
    repliedMessage := c.Message().ReplyTo
    requestBody := createRequestBody(repliedMessage.Text)
    responseBody, err := sendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    content, err := parseResponse(responseBody)
    if err != nil {
        return err
    }

    return c.Reply(content, telebot.ModeHTML)
}

func createRequestBody(userInput string) []byte {
    requestBody := map[string]interface{}{
        "model": "meta-llama/Meta-Llama-3.1-70B-Instruct",
        "messages": []map[string]string{
            {"role": "system", "content": systemRole},
            {"role": "user", "content": userInput},
        },
    }

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        log.Fatalf("Error marshalling JSON: %v", err)
    }
    return jsonData
}

func sendRequest(jsonData []byte, apiKey string) ([]byte, error) {
    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}

func parseResponse(body []byte) (string, error) {
    var response Response
    if err := json.Unmarshal(body, &response); err != nil {
        return "", fmt.Errorf("error unmarshalling JSON: %v", err)
    }

    if len(response.Choices) > 0 {
        return response.Choices[0].Message.Content, nil
    }
    return "", fmt.Errorf("no choices found in the response")
}
