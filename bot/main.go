package main

import (
	log "github.com/sirupsen/logrus"
	"flag"
	"github.com/adamc295/roombot/common"
)

var (
	f_token string
)

func init() {
	flag.StringVar(&f_token, "token", "", "The token for the bot")
}

func main() {
	flag.Parse()
	
	log.SetFormatter(&log.TextFormatter {
		DisableTimestamp: true,
		ForceColors: false,
	})
	
	// No token? Get out!
	if f_token == "" {
		log.Error("No token specified. Use -token to set the bot token")
		return
	}
	
	log.Info("Starting RoomBot...")
	
	err := common.Init()
	if err != nil {
		log.withError(err).Fatal("Failed on initialize")
	}
}