package msa

type ActivityType string

const (
	ActivityTypePublish ActivityType = "PUBLISH"
	ActivityTypeEdit    ActivityType = "EDIT"
	ActivityTypeDelete  ActivityType = "DELETE"
	ActivityTypeLike    ActivityType = "LIKE"
	ActivityTypeDislike ActivityType = "DISLIKE"
	ActivityTypeJoin    ActivityType = "JOIN"
	ActivityTypeLeave   ActivityType = "LEAVE"
	ActivityTypeMention ActivityType = "MENTION"
	ActivityTypeAdd     ActivityType = "ADD"
	ActivityTypeRemove  ActivityType = "REMOVE"
)
