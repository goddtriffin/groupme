package groupme

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Message is a GroupMe message.
type Message struct {
	ID         string `json:"id"`
	SourceGUID string `json:"source_guid"`
	UserID     string `json:"user_id"`
	GroupID    string `json:"group_id"`
	SenderID   string `json:"sender_id"`

	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Text      string `json:"text"`

	SenderType string `json:"sender_type"`
	Platform   string `json:"platform"`

	CreatedAt int  `json:"created_at"`
	System    bool `json:"system"`

	FavoritedBy []string     `json:"favorited_by"`
	Attachments []Attachment `json:"attachments"`
	Event       Event        `json:"event"`
}

// GetMessagesResponse is a the HTTP response from GetMessages (`GET /groups/:group_id/messages`).
type GetMessagesResponse struct {
	Count    int        `json:"count"`
	Messages []*Message `json:"messages"`
}

// GetMessages retrieves messages for a group.
func (c *Client) GetMessages(groupID string, limit string, beforeID, sinceID, afterID string) (GetMessagesResponse, error) {
	// build query params
	values := url.Values{}
	values.Add("token", c.AccessToken)
	if limit != "" {
		values.Add("limit", limit)
	}
	if beforeID != "" {
		values.Add("before_id", beforeID)
	}
	if sinceID != "" {
		values.Add("since_id", sinceID)
	}
	if afterID != "" {
		values.Add("after_id", afterID)
	}
	params := values.Encode()

	// generate URL for request
	URL, err := createURL(c.BaseURL, fmt.Sprintf("/groups/%s/messages", groupID), params)
	if err != nil {
		return GetMessagesResponse{}, err
	}

	// send request, read body
	resp, err := http.Get(URL)
	if err != nil {
		return GetMessagesResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return GetMessagesResponse{}, err
	}

	// exit early on error
	if resp.StatusCode == http.StatusNotModified {
		return GetMessagesResponse{}, ErrNotModified
	}

	// parse response
	var messages struct {
		Response GetMessagesResponse `json:"response"`
		Meta     Meta                `json:"meta"`
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return GetMessagesResponse{}, err
	}

	// exit early on error
	if messages.Meta.Code != http.StatusOK {
		return GetMessagesResponse{}, fmt.Errorf("%d: %s", messages.Meta.Code, fmt.Sprintf("%+v", messages.Meta.Errors))
	}

	return messages.Response, nil
}

// AllMessages retrieves all messages from a particular group.
func (c *Client) AllMessages(groupID string) ([]*Message, error) {
	var history []*Message

	var beforeID string
	for {
		messages, err := c.GetMessages(groupID, "100", beforeID, "", "")
		if err != nil {
			if errors.Is(err, ErrNotModified) {
				break
			}
			return nil, err
		}
		beforeID = messages.Messages[len(messages.Messages)-1].ID

		history = append(history, messages.Messages...)
	}

	return history, nil
}
