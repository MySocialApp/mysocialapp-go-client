package msa

type GroupMemberAccessControl string

const (
	GroupMemberAccessControlPublic         GroupMemberAccessControl = "PUBLIC"
	GroupMemberAccessControlFriendOfFriend GroupMemberAccessControl = "FRIEND_OF_FRIEND"
)
