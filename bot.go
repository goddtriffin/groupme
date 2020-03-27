package groupme

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

// Bot is a GroupMe Bot.
type Bot struct {
	BaseURL   string
	ID        string
	GroupID   string
	GroupName string
	AvatarURL string
}

// BotPost is a message from a Bot.
type BotPost struct {
	BotID       string       `json:"bot_id"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// NewBot returns a new GroupMe Bot.
func NewBot(baseURL, ID, groupID, groupName, avatarURL string) Bot {
	return Bot{
		BaseURL:   baseURL,
		ID:        ID,
		GroupID:   groupID,
		GroupName: groupName,
		AvatarURL: avatarURL,
	}
}

// Post posts a message.
func (b *Bot) Post(message string, attachments []Attachment) error {
	// generate URL for request
	URL, err := createURL(b.BaseURL, "/bots/post", "")
	if err != nil {
		return err
	}

	// chunk message down to lengths of 1000 or less
	for _, buf := range b.getBufferedMessage(message, "\n") {
		post := BotPost{
			BotID:       b.ID,
			Text:        buf,
			Attachments: attachments,
		}

		jsonStr, err := json.Marshal(post)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusAccepted {
			return parseError(resp.StatusCode, resp.Status)
		}
	}

	return nil
}

// getBufferedMessage returns a list of strings no bigger than
// what is allowed to be sent as a GroupMe Bot message (length: 1000).
func (b *Bot) getBufferedMessage(s, sep string) []string {
	list := []string{}
	var strBuilder string

	split := strings.Split(s, sep)
	for _, part := range split {
		if len(strBuilder)+len(part)+len(sep) <= 1000 {
			strBuilder += part + sep
		} else {
			list = append(list, strings.TrimSpace(strBuilder))
			strBuilder = part + sep
		}
	}

	if len(strBuilder) > 0 {
		list = append(list, strings.TrimSpace(strBuilder))
	}

	return list
}
