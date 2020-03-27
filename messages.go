package groupme

// IndexMessagesResponse is a `GET /groups/:group_id/messages` response.
type IndexMessagesResponse struct {
	Count    int       `json:"count"`
	Messages []Message `json:"messages"`
}

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
