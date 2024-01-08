package groupme

// Event types.
const (
	MemberAddedEventType           = "membership.announce.added"
	MemberRemovedEventType         = "membership.notifications.removed"
	MemberExitedEventType          = "membership.notifications.exited"
	MemberRejoinedEventType        = "membership.announce.rejoined"
	MemberNicknameChangedEventType = "membership.nickname_changed"
	MemberAvatarChangedEventType   = "membership.avatar_changed"

	GroupTypeChangeEventType   = "group.type_change"
	GroupAvatarChangeEventType = "group.avatar_change"
	GroupTopicChangeEventType  = "group.topic_change"

	PollCreatedEventType  = "poll.created"
	PollReminderEventType = "poll.reminder"
	PollFinishedEventType = "poll.finished"

	CalendarCreatedEventType         = "calendar.event.created"
	CalendarStartingEventType        = "calendar.event.starting"
	CalendarMemberGoinEventType      = "calendar.event.user.going"
	CalendarMemberUndecidedEventType = "calendar.event.user.undecided"
	CalendarMemberNotGoingEventType  = "calendar.event.user.not_going"
)

// EventData keys.
const (
	AdderUserKey   = "adder_user"
	AddedUsersKey  = "added_users"
	RemoverUserKey = "remover_user"
	RemovedUserKey = "removed_user"
)

// Event is an event tied to a GroupMe message.
type Event struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// UserEventData is a possible Event entry.
type UserEventData struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

// UsersEventData is a slice of UserEventDatas.
type UsersEventData []UserEventData

// Exists returns whether Event exists or not.
func (e *Event) Exists() bool {
	return e.Type == ""
}

// ParseUserEventData parses a UserEventData from an interface if possible.
func ParseUserEventData(i interface{}) (UserEventData, bool) {
	


	d := UserEventData{}

	m, ok := i.(map[string]interface{})
	if !ok {
		return d, ok
	}

	// parse ID
	tmp, ok := m["id"]
	if !ok {
		return d, ok
	}
	id, ok := tmp.(float64)
	if !ok {
		return d, ok
	}
	d.ID = int(id)

	// parse Nickname
	tmp, ok = m["nickname"]
	if !ok {
		return d, ok
	}
	nickname, ok := tmp.(string)
	if !ok {
		return d, ok
	}
	d.Nickname = nickname

	return d, true
}

// ParseUsersEventData parses a UserEventData from an interface if possible.
func ParseUsersEventData(i interface{}) (UsersEventData, bool) {
	ds := UsersEventData{}

	list, ok := i.([]interface{})
	if !ok {
		return ds, ok
	}

	for _, tmp := range list {
		d, ok := ParseUserEventData(tmp)
		if !ok {
			return ds, ok
		}

		ds = append(ds, d)
	}

	return ds, true
}
