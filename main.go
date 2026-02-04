package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type Application struct {
	client *tbot.Client
}

var (
	app   Application
	token string
	bot   *tbot.Server
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN environment variable is required")
	}

	log.Println("Starting bot...")

	bot = tbot.New(token)
	app.client = bot.Client()

	bot.HandleMessage(`https?://(?:www\.)?(?:twitter|x)\.com/[^/]+/status/\d+`, MessageHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := bot.Start(); err != nil {
			log.Printf("Bot error: %v", err)
		}
		stop()
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")
}
