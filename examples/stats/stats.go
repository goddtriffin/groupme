package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/MagnusFrater/groupme"
)

type stats struct {
	Members            map[string]*member
	WordFrequency      map[string]*word
	CharacterFrequency map[rune]*character
}

type member struct {
	ID              string
	Name            string
	PopularityScore int // how often did others upvote them
	SimpScore       int // how many times did they upvote someone else
	NarcissistScore int // how many times did they upvote themselves
	NumMessages     int // how many messages did they send
}

type character struct {
	R         rune
	Frequency int
}

type word struct {
	Text      string
	Frequency int
}

func newStats() stats {
	return stats{
		Members:            make(map[string]*member),
		CharacterFrequency: make(map[rune]*character),
		WordFrequency:      make(map[string]*word),
	}
}

func (s *stats) parseMessages(messages []groupme.Message) {
	for _, message := range messages {
		// parse numMessage and popularity
		s.incNumMessages(message.UserID, message.Name)
		s.incPopularity(message.UserID, message.Name, len(message.FavoritedBy))

		// parse narcissists and simps
		for _, userID := range message.FavoritedBy {
			if userID == message.UserID {
				s.incNarcissist(message.UserID, message.Name)
			} else {
				s.incSimp(userID, "")
			}
		}

		// parse word frequency
		for _, text := range strings.Fields(message.Text) {
			s.incWord(text)

			// parse character frequency
			runes := []rune(text)
			for _, r := range runes {
				s.incCharacter(r)
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

func (s *stats) addCharacter(r rune) {
	if _, ok := s.CharacterFrequency[r]; !ok {
		s.CharacterFrequency[r] = &character{R: r}
	}
}

func (s *stats) addWord(text string) {
	if _, ok := s.WordFrequency[text]; !ok {
		s.WordFrequency[text] = &word{Text: text}
	}
}

func (s *stats) incNumMessages(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].NumMessages++
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

func (s *stats) incWord(text string) {
	s.addWord(text)

	s.WordFrequency[text].Frequency++
}

func (s *stats) incCharacter(r rune) {
	s.addCharacter(r)

	s.CharacterFrequency[r].Frequency++
}

func (s *stats) topOfThePops(limit int) []*member {
	sorted := []*member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].PopularityScore > sorted[j].PopularityScore })

	top := []*member{}
	for i := 0; i < limit && i < len(sorted); i++ {
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
	for i := 0; i < limit && i < len(sorted); i++ {
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
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) topPosters(limit int) []*member {
	sorted := []*member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].NumMessages > sorted[j].NumMessages })

	top := []*member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) topWords(limit int) []*word {
	sorted := []*word{}

	for _, w := range s.WordFrequency {
		sorted = append(sorted, w)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Frequency > sorted[j].Frequency })

	top := []*word{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) topCharacters(limit int) []*character {
	sorted := []*character{}

	for _, c := range s.CharacterFrequency {
		sorted = append(sorted, c)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Frequency > sorted[j].Frequency })

	top := []*character{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

func (s *stats) sprintTopOfThePops(limit int) string {
	str := "Top of the Pops\n(who has the most upvotes)\n==========\n"

	topPopulars := s.topOfThePops(limit)
	for i, member := range topPopulars {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.PopularityScore)

		// don't put newline after last ranking
		if i < len(topPopulars)-1 {
			str += "\n"
		}
	}

	return str
}

func (s *stats) sprintTopOfTheSimps(limit int) string {
	str := "Top of the Simps\n(who upvoted other people the most)\n==========\n"

	topSimps := s.topOfTheSimps(limit)
	for i, member := range topSimps {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.SimpScore)

		// don't put newline after last ranking
		if i < len(topSimps)-1 {
			str += "\n"
		}
	}

	return str
}

func (s *stats) sprintTopOfTheNarcissists(limit int) string {
	str := "Top of the Narcissists\n(who upvoted themselves the most)\n==========\n"

	topNarcissists := s.topOfTheNarcissists(limit)
	for i, member := range topNarcissists {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.NarcissistScore)

		// don't put newline after last ranking
		if i < len(topNarcissists)-1 {
			str += "\n"
		}
	}

	return str
}

func (s *stats) sprintTopPoster(limit int) string {
	str := "Top Poster\n(who posted the most)\n==========\n"

	topPosters := s.topPosters(limit)
	for i, member := range topPosters {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.NumMessages)

		// don't put newline after last ranking
		if i < len(topPosters)-1 {
			str += "\n"
		}
	}

	return str
}

func (s *stats) sprintTopWords(limit int) string {
	str := "Top Words\n==========\n"

	topWords := s.topWords(limit)
	for i, w := range topWords {
		str += fmt.Sprintf("%d) %s: %d", i+1, w.Text, w.Frequency)

		// don't put newline after last ranking
		if i < len(topWords)-1 {
			str += "\n"
		}
	}

	return str
}

func (s *stats) sprintTopCharacters(limit int) string {
	str := "\nTop Characters\n==========\n"

	topCharacters := s.topCharacters(limit)
	for i, c := range topCharacters {
		str += fmt.Sprintf("%d) %s: %d", i+1, string(c.R), c.Frequency)

		// don't put newline after last ranking
		if i < len(topCharacters)-1 {
			str += "\n"
		}
	}

	return str
}

func (s *stats) sprintFrequencyAnalysis(limit int) string {
	str := "KOWALSKI, ANALYSIS!\n\n"

	str += fmt.Sprintf("%s\n%s",
		s.sprintTopWords(limit),
		s.sprintTopCharacters(limit))

	return str
}

func (s *stats) printAllMembers() {
	for _, member := range s.Members {
		fmt.Printf("%s (%s)\n==========\nPopularity Score: %d\nSimp Score: %d\nNarcissist Score: %d\n\n",
			member.Name, member.ID, member.PopularityScore, member.SimpScore, member.NarcissistScore)
	}
}
