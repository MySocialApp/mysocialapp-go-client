package msa

type AccessControl string

const (
	AccessControlPublic  AccessControl = "PUBLIC"
	AccessControlPrivate AccessControl = "PRIVATE"
	AccessControlFriend  AccessControl = "FRIEND"
)
