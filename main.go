package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	token, _ := os.LookupEnv("BOT_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	rand.Seed(time.Now().Unix())

	if m.Content == "/real-quote" {
		randomQuote := RealLifeQuotes[rand.Intn(len(RealLifeQuotes))]

		if _, err := s.ChannelMessageSend(m.ChannelID, randomQuote); err != nil {
			fmt.Println(err)
		}
	}

	if m.Content == "/movie-quote" {
		randomQuote := MovieQuotes[rand.Intn(len(MovieQuotes))]

		if _, err := s.ChannelMessageSend(m.ChannelID, randomQuote); err != nil {
			fmt.Println(err)
		}
	}

	type GiphyResp struct {
		Data []struct {
			EmbedUrl string `json:"embed_url"`
		} `json:"data"`
	}

	if m.Content == "/seagal-gif" {
		giphyApiKey, _ := os.LookupEnv("GIPHY_API_KEY")
		resp, err := http.Get(fmt.Sprintf("https://api.giphy.com/v1/gifs/search?api_key=%s&q=steven+seagal&limit=10&offset=0&lang=en", giphyApiKey))

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}
		var result GiphyResp

		if err := json.Unmarshal(body, &result); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Looks like giphy might be **Out of Reach**")
		}

		randomGif := result.Data[rand.Intn(len(result.Data))]

		if _, err := s.ChannelMessageSend(m.ChannelID, randomGif.EmbedUrl); err != nil {
			fmt.Println(err)
		}
	}
}
