package main

import (
    "log"
    "net/http"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "telegram-bot-go/internal/handler"
)

func main() {
    token := os.Getenv("1325102451:AAF4_JiwGDaQ9f8GyyRphl4UWssCuTgI9TA")
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
