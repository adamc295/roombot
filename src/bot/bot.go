package bot

import (
	log "github.com/sirupsen/logrus"
	"sync"
	
	"github.com/jonas747/dshardmanager"
	"github.com/jonas747/discordgo"
)

var (
	ShardManager *dshardmanager.Manager
	
	totalshardcount int
	
	processshards []int
)

func Run() {
	log.Info("Running the bot...")
	
	ShardManager = dshardmanager.New("Bot <TOKEN GOES HERE COULDNT BE BOTHERED TO MAKE A CONFIG SORRY>")
	ShardManager.Name = "RoomBot"
	
	ShardManager.SessionFunc = func(token string) (session *discordgo.Session, err error) {
		session, err = discordgo.New(token)
		if err != nil {
			return
		}
		
		session.AddHandler(OnMessageSend)
		
		session.StateEnabled = true
		session.LogLevel = discordgo.LogInformational
		session.SyncEvents = true
		
		session.State.Ready = discordgo.Ready {
			User: &discordgo.SelfUser {
				User: nil,
			},
		}
		
		return
	}
	
	shardcount, err := ShardManager.GetRecommendedCount()
	if err != nil {
		panic("Failed to get the recommended amount of shards: " + err.Error())
	}
	
	totalshardcount = shardcount
	processshards = make([]int, totalshardcount)
	for i := 0; i < totalshardcount; i++ {
		processshards[i] = i
	}
	
	// If you want to kill your bandwidth, set this to 100.
	ShardManager.SetNumShards(10)
	
	go ShardManager.Start()
}

func OnMessageSend(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		log.Info("Message was created by bot.")
		return
	}
	
	if message.Content == "roombot speak" {
		session.ChannelMessageSend(message.ChannelID, "Hello! I'm @RoomBot!")
	}
}

func Stop(wg *sync.WaitGroup) {
	ShardManager.StopAll()
	wg.Done()
}