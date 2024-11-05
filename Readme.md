
# English Teacher Bot

This project is a Telegram bot that acts as an English teacher, providing sentence corrections and explanations. Users can reply to a message with the command `/check` to receive feedback on their text.

## Features

- Corrects sentences and provides a revised version.
- Lists corrected errors in bullet points.
- Explains corrections in English.
  
## Requirements

- Go (1.16 or later)
- A Telegram Bot Token
- An API key for DeepInfra's OpenAI service
- An `.env` file for storing sensitive information

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/alirezadp10/english-teacher-bot.git
   cd english-teacher-bot
   ```

2. **Create a `.env` file** in the project root and add your bot token and API key:

   ```plaintext
   BOT_TOKEN=your_telegram_bot_token
   API_KEY=your_deepinfra_api_key
   ```

3. **Install dependencies:**

   ```bash
   go get -u github.com/joho/godotenv
   go get -u gopkg.in/telebot.v4
   ```

4. **Run the bot:**

   ```bash
   go run main.go
   ```

## Usage

To use the bot, follow these steps:

1. Open a chat with your bot on Telegram.
2. Send any message that you want to check.
3. Reply to that message with the command:

   ```plaintext
   /check
   ```

The bot will respond with the corrected version of your sentence, along with a list of corrections and their explanations.

## Contributing

Contributions are welcome! If you have suggestions or improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
