package groupme

// A polymorphic list of Message attachment types.
const (
	ImageAttachment    = "image"    // contains: Type, URL
	LocationAttachment = "location" // contains: Type, Name, Lat, Lng
	SplitAttachment    = "split"    // contains: Type, Token
	EmojiAttachment    = "emoji"    // contains: Type, Placeholder, Charmap
	MentionsAttachment = "mentions" // contains: Type, UserIDs, Loci
)

// Attachment is a GroupMe attachment.
type Attachment struct {
	// shared
	Type string `json:"type"`

	// Image
	URL string `json:"url"`

	// Location
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`

	// Split
	Token string `json:"token"`

	// Emoji
	Placeholder string  `json:"placeholder"`
	Charmap     [][]int `json:"charmap"`

	// Mentions
	UserIDs []string `json:"user_ids"`
	Loci    [][]int  `json:"loci"`
}
