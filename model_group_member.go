package msa

type GroupMember struct {
	CreatedDate    string      `json:"created_date"`
	UpdatedDate    string      `json:"updated_date"`
	Group          Group       `json:"group"`
	User           User        `json:"user"`
	PreviousStatus *GroupStatus `json:"previous_status"`
	Status         *GroupStatus `json:"status"`
}
