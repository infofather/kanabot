package handler

import (
    "encoding/json"
    "net/http"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WebhookHandler(bot *tgbotapi.BotAPI) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var update tgbotapi.Update
        if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        if update.Message != nil {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я бот на Go!")
            bot.Send(msg)
        }

        w.WriteHeader(http.StatusOK)
    }
}
