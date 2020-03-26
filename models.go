package groupme

// Meta is a
type Meta struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}

// IndexMessagesResponse is a GET /groups/:group_id/messages response.
type IndexMessagesResponse struct {
	Count    int       `json:"count"`
	Messages []Message `json:"messages"`
}

// Message is a GroupMe message.
type Message struct {
	ID          string       `json:"id"`
	SourceGUID  string       `json:"source_guid"`
	CreatedAt   int          `json:"created_at"`
	UserID      string       `json:"user_id"`
	GroupID     string       `json:"group_id"`
	Name        string       `json:"name"`
	AvatarURL   string       `json:"avatar_url"`
	Text        string       `json:"text"`
	System      bool         `json:"system"`
	FavoritedBy []string     `json:"favorited_by"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment is a GroupMe attachment.
type Attachment struct {
	Type        string  `json:"type"`
	URL         string  `json:"url"`
	Lat         string  `json:"lat"`
	Lng         string  `json:"lng"`
	Name        string  `json:"name"`
	Token       string  `json:"token"`
	Placeholder string  `json:"placeholder"`
	Charmap     [][]int `json:"charmap"`
}

// BotPost . . . TODO
type BotPost struct {
	BotID string `json:"bot_id"`
	Text  string `json:"text"`
}
