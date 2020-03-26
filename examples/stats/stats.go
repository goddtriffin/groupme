package main

import (
	"fmt"
	"sort"

	"github.com/MagnusFrater/groupme"
)

type stats struct {
	Members     map[string]*member
	TopMessages []groupme.Message
}

type member struct {
	ID              string
	Name            string
	PopularityScore int // how often did others upvote them
	SimpScore       int // how many times did they upvote someone else
	NarcissistScore int // how many times did they upvote themselves
}

func newStats(limitTopMessages int) stats {
	return stats{
		Members:     make(map[string]*member),
		TopMessages: make([]groupme.Message, 0, limitTopMessages),
	}
}

func (s *stats) parseMessages(messages []groupme.Message) {
	for _, message := range messages {
		s.incPopularity(message.UserID, message.Name, len(message.FavoritedBy))

		for _, userID := range message.FavoritedBy {
			if userID == message.UserID {
				s.incNarcissist(message.UserID, message.Name)
			} else {
				s.incSimp(userID, "")
			}
		}
	}
}

func (s *stats) addMember(userID, name string) {
	if m, ok := s.Members[userID]; !ok {
		s.Members[userID] = &member{
			ID:   userID,
			Name: name,
		}
	} else {
		if m.Name == "" {
			m.Name = name
		}
	}
}

func (s *stats) incPopularity(userID, name string, inc int) {
	s.addMember(userID, name)

	s.Members[userID].PopularityScore += inc
}

func (s *stats) incSimp(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].SimpScore++
}

func (s *stats) incNarcissist(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].NarcissistScore++
}

func (s *stats) topOfThePops(limit int) []*member {
	sorted := []*member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].PopularityScore > sorted[j].PopularityScore })

	top := []*member{}
	for i := 0; i < limit; i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) topOfTheSimps(limit int) []*member {
	sorted := []*member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].SimpScore > sorted[j].SimpScore })

	top := []*member{}
	for i := 0; i < limit; i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) topOfTheNarcissists(limit int) []*member {
	sorted := []*member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].NarcissistScore > sorted[j].NarcissistScore })

	top := []*member{}
	for i := 0; i < limit; i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) sprintTopOfThePops(limit int) string {
	str := "Top of the Pops\n(who has the most upvotes)\n==========\n"

	for i, member := range s.topOfThePops(limit) {
		str += fmt.Sprintf("%d) %s: %d\n", i+1, member.Name, member.PopularityScore)
	}

	return str
}

func (s *stats) sprintTopOfTheSimps(limit int) string {
	str := "Top of the Simps\n(who upvoted other people the most)\n==========\n"

	for i, member := range s.topOfTheSimps(limit) {
		str += fmt.Sprintf("%d) %s: %d\n", i+1, member.Name, member.SimpScore)
	}

	return str
}

func (s *stats) sprintTopOfTheNarcissists(limit int) string {
	str := "Top of the Narcissists\n(who upvoted themselves the most)\n==========\n"

	for i, member := range s.topOfTheNarcissists(limit) {
		str += fmt.Sprintf("%d) %s: %d\n", i+1, member.Name, member.NarcissistScore)
	}

	return str
}

func (s *stats) printAllMembers() {
	for _, member := range s.Members {
		fmt.Printf("%s (%s)\n==========\nPopularity Score: %d\nSimp Score: %d\nNarcissist Score: %d\n\n",
			member.Name, member.ID, member.PopularityScore, member.SimpScore, member.NarcissistScore)
	}
}
