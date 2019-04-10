package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"time"
	"syscall"
	
	"github.com/adamc295/roombot/src/bot"
	"github.com/adamc295/roombot/src/web"
)

var (
	f_rune bool
	f_runb bool
	f_runw bool
)

func init() {
	// Flags and stuff
	flag.BoolVar(&f_runb, "bot", false, "Set this flag to run the bot")
	flag.BoolVar(&f_runw, "web", false, "Set this flag to run the webserver")
	
	flag.BoolVar(&f_rune, "everything", false, "Set this flag to run the bot and webserver")
}

func main() {
	// Get our flags
	flag.Parse()
	
	// Formatting for the logs
	log.SetFormatter(&log.TextFormatter {
		DisableTimestamp: true,
		ForceColors: true,
	})
	
	// Nothing was set
	if !f_runb && !f_runw && !f_rune {
		log.Error("Nothing was set to run! See --help for more")
		return
	}
	
	log.Info("Starting RoomBot...")
	
	// Run the web server
	if f_runw || f_rune {
		go web.Run()
	}
	
	// Run the bot
	if f_runb || f_rune {
		bot.Run()
	}
	
	listenForSignal()
}

func listenForSignal() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	sig := <-c
	log.Info("Shutting down: ", sig.String())
	
	shoudwait := false
	wg := new(sync.WaitGroup)
	
	if f_runb || f_rune {
		wg.Add(1)
		go bot.Stop(wg)
		shoudwait = true
	}
	
	if f_runw || f_rune {
		web.Stop()
		time.Sleep(time.Second)
	}
	
	if shoudwait {
		log.Info("Waiting to shut down...")
		wg.Wait()
	}
	
	log.Info("Sleeping for a second...")
	time.Sleep(time.Second)
	
	log.Info("Goodnight...")
	os.Exit(0)
}