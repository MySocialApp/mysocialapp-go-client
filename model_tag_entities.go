package msa

type TagEntities struct {
	UserMentionTags []UserMentionTag `json:"user_mention_tags"`
	UrlTags         []UrlTag         `json:"url_tags"`
	HashTags        []HashTag        `json:"hash_tags"`
}

func (t *TagEntities) GetAllTagEntities() []TagEntity {
	list := []TagEntity{}
	if t.UserMentionTags != nil {
		for _, t := range t.UserMentionTags {
			list = append(list, t)
		}
	}
	if t.UrlTags != nil {
		for _, t := range t.UrlTags {
			list = append(list, t)
		}
	}
	if t.HashTags != nil {
		for _, t := range t.HashTags {
			list = append(list, &t)
		}
	}
}
