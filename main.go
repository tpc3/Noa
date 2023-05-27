package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"

	"github.com/bluele/mecab-golang"
	"github.com/robfig/cron/v3"
	"github.com/tpc3/Noa/lib/config"
	"github.com/tpc3/Noa/lib/markov"
	misskeyapi "github.com/tpc3/Noa/lib/misskeyApi"
)

func main() {
	botID := misskeyapi.MisskeyGetuserID(config.Loadconfig.Misskey.Token)
	c := cron.New()
	c.AddFunc("@hourly", func() { runMarkov(botID) })
	c.Start()
	log.Print("Noa is now running!")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func runMarkov(botID string) {
	var text string
	markovBlock := [][]string{}
	notes, err := misskeyapi.MisskeyGetnotesRequest(config.Loadconfig.Misskey.Token, botID)
	if err != nil {
		return
	}
	m, err := mecab.New("-Owakati")
	if err != nil {
		log.Fatal("New mecab error: ", err)
	}
	defer m.Destroy()

	for i := 0; i < 10; i++ {
		for _, note := range notes {
			_data := markov.ParseToNode(m, note)
			elems := markov.GetMarkovBlocks(_data)
			markovBlock = append(markovBlock, elems...)
		}

		noteElemset := markov.MarkovChainExec(markovBlock)
		text = markov.TextGenerate(noteElemset)

		ok := true
		for _, v := range config.Loadconfig.TextBlacklist {
			check := regexp.MustCompile(v)
			if check.MatchString(text) {
				ok = false
			}
		}

		if ok {
			break
		}
	}

	err = misskeyapi.MisskeySendnotesRequest(config.Loadconfig.Misskey.Token, text)
	if err != nil {
		log.Print(text)
		return
	}
}
