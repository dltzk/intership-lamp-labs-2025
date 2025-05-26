package handlers

import (
	"fmt"
	telebot "gopkg.in/telebot.v4"
	redis "github.com/redis/go-redis/v9"
	"context"
	"time"
	"log"
)

func Ping(rdb *redis.Client) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		userID := c.Sender().ID
        key := fmt.Sprintf("user:%d:pings", userID)
		_, err := rdb.Incr(context.Background(), key).Result()
		if err != nil {
			log.Printf("Пользователю %d не удалось вызвать /ping.", userID)
            return c.Send("Не удалось обработать запрос.")
        }
		log.Printf("Пользователь %d вызвал /ping.", userID)
		return c.Send("pong")
	}
}

func PingMyPings(rdb *redis.Client) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		userID := c.Sender().ID
        key := fmt.Sprintf("user:%d:pings", userID)
		count, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			log.Printf("Пользователю %d не удалось вызвать /mypings.", userID)
            return c.Send("Не удалось обработать запрос.")
        }
		log.Printf("Пользователь %d вызвал /mypings.", userID)
		return c.Send(fmt.Sprintf("Вы отправили %s пингов.", count))
	}
}

func TellMeTime() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		log.Printf("Пользователь %d вызвал /time.", c.Sender().ID)
		return c.Send(time.Now().Format(time.DateTime))
	}
}
