package groupme

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Bot is a GroupMe Bot.
type Bot struct {
	BaseURL   string
	ID        string
	GroupID   string
	GroupName string
	AvatarURL string
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

// Post . . . TODO
func (b *Bot) Post(message string) error {
	// generate URL for request
	URL, err := createURL(b.BaseURL, "/bots/post", "")
	if err != nil {
		return err
	}

	post := BotPost{
		BotID: b.ID,
		Text:  message,
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

	return nil
}
