package groupme

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is a GroupMe API client.
type Client struct {
	BaseURL     string
	AccessToken string
}

// NewClient returns a new GroupMe API client.
func NewClient(baseURL, accessToken string) Client {
	return Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
	}
}

// IndexMessages retrieves messages for a group.
func (c *Client) IndexMessages(groupID string, limit string, beforeID, sinceID, afterID string) (IndexMessagesResponse, error) {
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
		return IndexMessagesResponse{}, err
	}

	// send request, read body
	resp, err := http.Get(URL)
	if err != nil {
		return IndexMessagesResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return IndexMessagesResponse{}, err
	}

	// exit early on error
	if resp.StatusCode == http.StatusNotModified {
		return IndexMessagesResponse{}, ErrNotModified
	}

	// parse response
	var messages struct {
		Response IndexMessagesResponse `json:"response"`
		Meta     Meta                  `json:"meta"`
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return IndexMessagesResponse{}, err
	}

	// exit early on error
	if messages.Meta.Code != http.StatusOK {
		return IndexMessagesResponse{}, fmt.Errorf("%d: %s", messages.Meta.Code, fmt.Sprintf("%+v", messages.Meta.Errors))
	}

	return messages.Response, nil
}

// AllMessages retrieves all messages from a particular group.
func (c *Client) AllMessages(groupID string) ([]Message, error) {
	var history []Message

	var beforeID string
	for {
		messages, err := c.IndexMessages(groupID, "100", beforeID, "", "")
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
