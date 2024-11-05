
# Project Name

This project is a Telegram bot that interacts with users, processes messages, and performs various database operations to save message details and user information. The bot also communicates with an external API to enhance its functionalities.

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- PostgreSQL
- A Telegram bot token (obtainable from the [BotFather](https://core.telegram.org/bots#botfather))
- API key for the external service (replace `API_URL` in `api/client.go` if necessary)

### Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/alirezadp10/english-editor-bot.git
   cd projectname
   ```

2. **Configure Environment Variables**:
   Create a `.env` file in the root directory and add the necessary environment variables:

   ```plaintext
   BOT_TOKEN=your_telegram_bot_token
   API_KEY=your_api_key
   POSTGRES_USER=your_postgres_username
   POSTGRES_PASSWORD=your_postgres_password
   POSTGRES_DB=your_postgres_db
   POSTGRES_PORT=your_postgres_port
   ```

3. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the Bot**:
   ```bash
   go run main.go
   ```

## Code Structure Details

### main.go

The main entry point initializes the bot, connects to the database, and starts the bot with configured handlers.

### bot/

- **bot.go**: Initializes the bot and defines its settings.
- **handlers.go**: Contains bot command handlers and message processing functions.

### database/

- **connection.go**: Manages the PostgreSQL database connection.
- **models.go**: Defines models for the database (User, Chat, Message, etc.).
- **operations.go**: Includes functions to perform operations like saving users, chats, and messages.

### api/

- **client.go**: Handles the creation and sending of requests to the external API.
- **parser.go**: Parses responses from the external API.

### config/

- **config.go**: Loads environment variables from the `.env` file.

## Usage

- After starting the bot, you can interact with it on Telegram by sending messages.
- Use `/check` command to trigger the bot's message processing and API interaction.

## License

This project is licensed under the MIT License. See `LICENSE` for details.
