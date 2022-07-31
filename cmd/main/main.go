package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"telebot/chuckNorris"
	"telebot/config"
	"telebot/logger"
	"telebot/weather"
	"telebot/youtube"
	"time"
)

func main() {
	logger := logger.NewBuiltinLogger()
	cfg := config.GetConfig()

	pref := tele.Settings{
		Token:  cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/yt", func(c tele.Context) error {
		sendText := "Твой трек не наиден"
		trackName := c.Message().Payload
		name, err := youtube.FindByTrack(trackName)
		if err != nil {
			return c.Send(sendText)
		}
		sendText = name
		return c.Send(fmt.Sprintf("Вот твой трек: %s", sendText))
	})

	b.Handle("/ch", func(c tele.Context) error {
		joke, err := chuckNorris.GetJoke()
		if err != nil {
			return err
		}
		return c.Send(joke)
	})

	b.Handle("/wh", func(c tele.Context) error {

		city := c.Message().Payload
		weather, err := weather.FindWeatherCity(city)
		if err != nil {
			return c.Send(weather)
		}
		//sendWeatherMessage = weather
		return c.Send(fmt.Sprintf("Вот твоя погода: %s", weather))
	})

	logger.InfoLog.Println("Start application")
	b.Start()
}
