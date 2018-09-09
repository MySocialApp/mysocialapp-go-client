package msa

type GroupStatus string

const (
	GroupStatusWaitingForApproval    GroupStatus = "WAITING_FOR_APPROVAL"
	GroupStatusMember                GroupStatus = "MEMBER"
	GroupStatusWasMember             GroupStatus = "WAS_MEMBER"
	GroupStatusWasWaitingForApproval GroupStatus = "WAS_WAITING_FOR_APPROVAL"
)
