package msa

type FriendRequests struct {
	Incoming []User `json:"incoming"`
	Outgoing []User `json:"outgoing"`
}
