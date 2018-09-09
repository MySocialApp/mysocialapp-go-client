package msa

type EventMemberAccessControl string

const (
	EventMemberAccessControlPublic         EventMemberAccessControl = "PUBLIC"
	EventMemberAccessControlFriendOfFriend EventMemberAccessControl = "FRIEND_OF_FRIEND"
)
