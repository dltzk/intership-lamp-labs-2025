package main

import (
	telebot "gopkg.in/telebot.v4"
	handlers "telegram-bot/internal/delivery/handlers"
	redis "telegram-bot/internal/infrastructure/redisclient"
	"os"
	"time"
	"log"
)

func main() {
	rdb := redis.InitRedis()
	log.Println("Redis started")

	pref := telebot.Settings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, _ := telebot.NewBot(pref)
	log.Println("Bot created")

	bot.Handle("/ping", handlers.Ping(rdb))
	bot.Handle("/ping/my", handlers.PingMyPings(rdb))
	bot.Handle("/time", handlers.TellMeTime())

	log.Println("Bot started!")
	bot.Start()
}