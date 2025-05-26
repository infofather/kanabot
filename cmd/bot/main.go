package main

import (
    "log"
    "net/http"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "telegram-bot-go/internal/handler"
)

func main() {
    token := os.Getenv("1325102451:AAHIGKxaFfRV2xnLkLJ5KV2fI36pNH8K5Vk")
    if token == "" {
        log.Fatal("BOT_TOKEN is not set")
    }

    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Panic(err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    http.HandleFunc("/webhook", handler.WebhookHandler(bot))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Listening on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
