# VX That For Me

A Telegram bot that automatically converts Twitter/X links to fixupx.com links for better embed support.

## Features

- Automatically detects Twitter/X links in messages
- Converts them to fixupx.com links for better embed previews
- Deletes original messages to keep chats clean
- Notifies users when a link is fixed

## Prerequisites

- Go 1.21.8 or later
- A Telegram bot token (get one from [@BotFather](https://t.me/botfather))

## Installation

1. Clone the repository:
```bash
git clone https://github.com/guifrrs/vx-that-for-me.git
cd vx-that-for-me
```

2. Create a `.env` file:
```bash
cp .env.example .env
```

3. Add your Telegram bot token to `.env`:
```
TELEGRAM_TOKEN=your_bot_token_here
```

4. Build and run:
```bash
go build -o main .
./main
```

## Docker

Build and run with Docker:

```bash
docker build -t vx-that-for-me .
docker run -e TELEGRAM_TOKEN=your_token_here vx-that-for-me
```

## Usage

1. Add the bot to a Telegram group chat
2. Send any Twitter/X link (e.g., `https://x.com/user/status/123456`)
3. The bot will:
   - Send a friendly message mentioning you
   - Reply with the fixed link
   - Delete your original message

## Development

| Command | Description |
|---------|-------------|
| `go test ./...` | Run all tests |
| `go test -run TestName ./...` | Run specific test |
| `go fmt ./...` | Format code |
| `go vet ./...` | Check for issues |
| `go mod tidy` | Clean up dependencies |

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `TELEGRAM_TOKEN` | Your Telegram bot token from @BotFather | Yes |

## License

MIT
