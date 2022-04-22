package main

import (
	"encoding/json"
	"flag"
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

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
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

var realLifeQuotes = [19]string{
	"My people are my people. They love me and I love them. I would not be here without them.",
	"I am hoping that I can be known as a great writer and actor some day, rather than a sex symbol.",
	"Any great warrior is also a scholar, and a poet, and an artist",
	"Can I laugh in your face?",
	"I have no fear of death. More important, I don't fear life.",
	"Aikido is not merely about fighting and the development of the physical self but the perfection of the spiritual man at the same time. It has very harmonious movements, very beautiful to watch and beautiful for your body to feel.",
	"You can say that I lived in Asia for a long time and in Japan I became close to several CIA agents. And you could say that I became an adviser to several CIA agents in the field and, through my friends in the CIA, met many powerful people and did special works and special favors.",
	"I'm a hip-hop fan.",
	"I was born in Detroit, in an all black neighborhood",
	"Zen Master Steven Seagal has become the Brand Ambassador of Bitcoiin2gen",
	"I have made a lot of mistakes.",
	"Most of the kids that I meet in the street are serious hardened criminals that I meet in the street, never had a mother and a father to love them, to protect them, to teach them right from wrong and lead them out of crime and gangs and stuff like that.",
	"I think we're living in a world where society is very difficult.",
	"My ancestors were Russian mongols.",
	"I look at both as one family.",
	"[Vladimir Putin is] one of the great living world leaders",
	"Looks like you're buying rounds tonight!",
	"See how it goes where I want it to? Right there!",
	"Notice how you're no longer doing this? Now you're doing _this_.",
}

var movieQuotes = [7]string{
	"Well I'm sorry to hear that. Because now, **I will snatch every motherfucker birthday**.",
	"I'd like to make something very clear. I don't have rage. I'm a happy guy. You see this face? This is a happy face. You all would be lucky to be as happy as I am!",
	"Women. You can't trust em.",
	"What kind of damn fool do you think I am? I'm still in prison for doing the same thing you're about to ask me to do again.",
	"Are you really that good?_Every once in a while._",
	"*sits*",
	"You gotta live to talk about old times.",
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	rand.Seed(time.Now().Unix())

	if m.Content == "/real-quote" {
		randomQuote := realLifeQuotes[rand.Intn(len(realLifeQuotes))]

		_, err := s.ChannelMessageSend(m.ChannelID, randomQuote)
		if err != nil {
			fmt.Println(err)
		}
	}

	if m.Content == "/movie-quote" {
		randomQuote := movieQuotes[rand.Intn(len(movieQuotes))]

		_, err := s.ChannelMessageSend(m.ChannelID, randomQuote)
		if err != nil {
			fmt.Println(err)
		}
	}

	type Analytics struct {
		EmbedUrl string `json:"embed_url"`
	}

	type GiphyResp struct {
		Data []Analytics `json:"data"`
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
		// why is a new json decoder a worse way than unmarshalling?
		json.Unmarshal(body, &result)

		gif := result.Data[rand.Intn(len(result.Data))]
		fmt.Println(gif.EmbedUrl)

		_, err = s.ChannelMessageSend(m.ChannelID, gif.EmbedUrl)

		if err != nil {
			fmt.Println(err)
		}
	}
}
