package msa

type EventMember struct {
	CreatedDate    string       `json:"created_date"`
	UpdatedDate    string       `json:"updated_date"`
	Event          *Event       `json:"event"`
	User           *User        `json:"user"`
	PreviousStatus *EventStatus `json:"previous_status"`
	Status         *EventStatus `json:"status"`
}
