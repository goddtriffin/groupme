package main

import (
	"flag"
	"log"
	"os"

	"github.com/MagnusFrater/groupme"
)

func main() {
	accessToken := flag.String("accessToken", "", "GroupMe API client access token")
	botID := flag.String("botID", "", "GroupMe Bot ID")
	groupID := flag.String("groupID", "", "Group ID")
	flag.Parse()

	if *accessToken == "" || *botID == "" || *groupID == "" {
		flag.Usage()
		os.Exit(1)
	}

	client := groupme.NewClient(groupme.V3BaseURL, *accessToken)
	bot := groupme.NewBot(groupme.V3BaseURL, *botID, *groupID, "", "")

	messages, err := client.AllMessages(bot.GroupID)
	if err != nil {
		log.Fatal(err)
	}

	stats := newStats(5)
	stats.parseMessages(messages)

	err = bot.Post(stats.sprintTopOfThePops(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.sprintTopOfTheSimps(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.sprintTopOfTheNarcissists(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.sprintTopPoster(5))
	if err != nil {
		log.Fatal(err)
	}
}
