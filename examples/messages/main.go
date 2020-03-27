package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/MagnusFrater/groupme"
)

func main() {
	accessToken := flag.String("accessToken", "", "GroupMe API client access token")
	botID := flag.String("botID", "", "GroupMe Bot ID")
	groupID := flag.String("groupID", "", "GroupMe Group ID")
	limit := flag.Int("limit", 5, "number of items to list")
	flag.Parse()

	if *accessToken == "" || *botID == "" || *groupID == "" {
		flag.Usage()
		os.Exit(1)
	}

	client := groupme.NewClient(groupme.V3BaseURL, *accessToken)
	bot := groupme.NewBot(groupme.V3BaseURL, *botID, *groupID, "", "")

	messages, err := client.GetMessages(bot.GroupID, strconv.Itoa(*limit), "", "", "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total messages: %d messages\n\n", messages.Count)

	fmt.Println("Most recent messages:")
	for i, message := range messages.Messages {
		fmt.Printf("%d) %s got %d favorites saying: \"%s\"\n", i+1, message.Name, len(message.FavoritedBy), message.Text)

		if message.Event.Exists() {
			switch message.Event.Type {
			case groupme.MemberAddedEventType:
				printMemberAddedEvent(message.Event)
			case groupme.MemberRemovedEventType:
				printMemberRemovedEvent(message.Event)
			case groupme.NicknameChangedEventType:
				printNicknameChangedEvent(message.Event)
			}
		}
	}
}

func printMemberAddedEvent(event groupme.Event) {
	var adderNickname, addedUsers string

	for _, i := range event.Data {
		userEventData, ok := groupme.ParseUserEventData(i)
		if ok {
			adderNickname = userEventData.Nickname
			continue
		}

		usersEventData, ok := groupme.ParseUsersEventData(i)
		if ok {
			for i, user := range usersEventData {
				addedUsers += user.Nickname

				if i < len(usersEventData)-1 {
					addedUsers += ","
				}
			}
		}
	}

	fmt.Printf("\tEvent: %s added %s\n", adderNickname, addedUsers)
}

func printMemberRemovedEvent(event groupme.Event) {
	var removerNickname, removedNickname string

	for key, i := range event.Data {
		userEventData, ok := groupme.ParseUserEventData(i)
		if ok {
			switch key {
			case groupme.RemoverUserKey:
				removerNickname = userEventData.Nickname
			case groupme.RemovedUserKey:
				removedNickname = userEventData.Nickname
			}
		}
	}

	fmt.Printf("\tEvent: %s removed %s\n", removerNickname, removedNickname)
}

func printNicknameChangedEvent(event groupme.Event) {
	var originalName, nicknameChangedTo string

	for _, i := range event.Data {
		userEventData, ok := groupme.ParseUserEventData(i)
		if ok {
			originalName = userEventData.Nickname
			continue
		}

		if str, ok := i.(string); ok {
			nicknameChangedTo = str
		}
	}

	fmt.Printf("\tEvent: %s changed their name to %s\n", originalName, nicknameChangedTo)
}
