package msa

type UserStat struct {
	Status *userStatStatus `json:"status"`
	Friend *userStatFriend `json:"friend"`
}

type userStatStatus struct {
	LastConnectionDate string          `json:"last_connection_date"`
	State              UserStatusState `json:"state"`
}

type userStatFriend struct {
	Total int `json:"total"`
}

type UserStatusState string

const (
	UserStatusStateDisabled     UserStatusState = "DISABLED"
	UserStatusStateUnknown      UserStatusState = "UNKNOWN"
	UserStatusStateConnected    UserStatusState = "CONNECTED"
	UserStatusStateAway         UserStatusState = "AWAY"
	UserStatusStateNotConnected UserStatusState = "NOT_CONNECTED"
)
