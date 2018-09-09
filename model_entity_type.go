package msa

type EntityType string

const (
	EntityTypeUser                EntityType = "USER"
	EntityTypeRide                EntityType = "RIDE"
	EntityTypeGroup               EntityType = "GROUP"
	EntityTypeEvent               EntityType = "EVENT"
	EntityTypePhoto               EntityType = "PHOTO"
	EntityTypePhotoAlbum          EntityType = "PHOTO_ALBUM"
	EntityTypeStatus              EntityType = "STATUS"
	EntityTypeConversation        EntityType = "CONVERSATION"
	EntityTypeTextWallMessage     EntityType = "TEXT_WALL_MESSAGE"
	EntityTypeStory               EntityType = "STORY"
	EntityTypeLocation            EntityType = "LOCATION"
	EntityTypeComment             EntityType = "COMMENT"
	EntityTypeLike                EntityType = "LIKE"
	EntityTypeUrlTag              EntityType = "URL_TAG"
	EntityTypeHashTag             EntityType = "HASH_TAG"
	EntityTypeUserMentionTag      EntityType = "USER_MENTION_TAG"
	EntityTypeConversationMessage EntityType = "CONVERSATION_MESSAGE"
)
