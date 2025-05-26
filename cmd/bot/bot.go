package main

import (
	telebot "gopkg.in/telebot.v4"
	handlers "telegram-bot/internal/delivery/handlers"
	redis "telegram-bot/internal/infrastructure/redisclient"
	"github.com/joho/godotenv"
	"time"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Ошибка при загрузке файла .env:", err)
	}
	rdb := redis.InitRedis()
	log.Println("Redis started!")

	token := os.Getenv("TELEGRAM_TOKEN")
    if token == "" {
        log.Fatal("TELEGRAM_TOKEN не установлен")
    }

	pref := telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
    log.Fatalf("Не удалось создать бота: %v", err)
	}

	log.Println("Bot created")

	bot.Handle("/ping", handlers.Ping(rdb))
	bot.Handle("/ping/my", handlers.PingMyPings(rdb))
	bot.Handle("/time", handlers.TellMeTime())

	log.Println("Bot started!")
	bot.Start()
}