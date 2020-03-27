package groupme

// Event types.
const (
	MembershipRemovedEvent = "membership.notifications.removed"
)

// EventData keys.
const (
	RemoverUserKey = "remover_user"
	RemovedUserKey = "removed_user"
)

// Event is an event tied to a GroupMe message.
type Event struct {
	Type string               `json:"type"`
	Data map[string]EventData `json:"data"`
}

// EventData is data tied to an Event.
type EventData struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}
