package msa

type Group struct {
	BaseWall
	Name                     string                    `json:"name"`
	Description              string                    `json:"description"`
	Open                     bool                      `json:"open"`
	Location                 *Location                 `json:"location"`
	ProfilePhoto             *Photo                    `json:"profile_photo"`
	ProfileCoverPhoto        *Photo                    `json:"profile_cover_photo"`
	Members                  []User                    `json:"members"`
	DistanceInMeters         *int                      `json:"distance_in_meters"`
	TotalMembers             *int                      `json:"total_members"`
	IsMember                 bool                      `json:"is_member"`
	GroupMemberAccessControl *GroupMemberAccessControl `json:"group_member_access_control"`
	CustomFields             []CustomField             `json:"custom_fields"`
}

func (g *Group) UpdateImage(image []byte) (*Photo, *RestError) {
	// TODO
	return nil, nil
}

func (g *Group) UpdateCoverImage(image []byte) (*Photo, *RestError) {
	// TODO
	return nil, nil
}

func (e *Group) ListNewsFeed(page int, size int) ([]Feed, *RestError) {
	// TODO
	return nil, nil
}

func (e *Group) AddNewsFeedMessage(feedPost *FeedPost) (*Feed, *RestError) {
	// TODO
	return nil, nil
}

func (e *Group) Update() *RestError {
	// TODO
	return nil
}

func (e *Group) Join() (*GroupMember, *RestError) {
	// TODO
	return nil, nil
}

func (e *Group) Quit() *RestError {
	// TODO
	return nil
}
