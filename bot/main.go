package main

// Hoo boy. Go.
import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jonas747/discordgo"
)

// The token is the username and password for the bot. If it gets out, then ya fucked.
var (
	Token string
)

func init() {
	// Get the token
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
}

func main() {
	// Login with the token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}
	resp, err := dg.GatewayBot()
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}
	dg.ShardCount = resp.Shards
	dg.ShardID = 0

	dg.LogLevel = discordgo.LogDebug

	// Add the handlers
	dg.AddHandler(messageCreate)

	// Open for business!
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Hooray! The bot is running. Press CTRL-C to quit.")
	sc := make(chan os.Signal, 1)
	// Basiclly, wait for us to die
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Die.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Received a message!")
	if m.Author.ID == s.State.User.ID {
		// Ignore, message was made by bot
		fmt.Println("Never mind, it was made by the bot.")
		return
	}

	if m.Content == "roombot speak" {
		fmt.Println("Speaking...")
		s.ChannelMessageSend(m.ChannelID, "Hello, World!")
	}
}
