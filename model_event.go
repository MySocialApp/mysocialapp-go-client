package msa

type Event struct {
	BaseWall
	Name                        string                    `json:"name"`
	Description                 string                    `json:"description"`
	IsCancelled                 *bool                     `json:"is_cancelled"`
	MaxSeats                    *int                      `json:"max_seats"`
	FreeSeats                   *int                      `json:"free_seats"`  // can be used, but only when in query param or direct /event/:id limited=false
	TakenSeats                  *int                      `json:"taken_seats"` // can be used, but only when in query param or direct /event/:id limited=false
	TotalMembers                *int                      `json:"total_members"`
	StartDate                   string                    `json:"start_date"`
	EndDate                     string                    `json:"end_date"`
	Members                     []EventMember             `json:"members"`
	EventMemberAccessControl    *EventMemberAccessControl `json:"event_member_access_control"`
	Location                    *Location                 `json:"location"`
	IsMember                    *bool                     `json:"is_member"`
	StaticMapsURL               *string                   `json:"static_maps_url"`
	DistanceInMeters            *int                      `json:"distance_in_meters"`
	ProfilePhoto                *Photo                    `json:"profile_photo"`
	ProfileCoverPhoto           *Photo                    `json:"profile_cover_photo"`
	IsAvailable                 *bool                     `json:"is_available"`
	RemainingSecondsBeforeStart *int                      `json:"remaining_seconds_before_start"`
	CustomFields                []CustomField             `json:"custom_fields"`
}

func (e *Event) UpdateImage(image []byte) (*Photo, *RestError) {
	// TODO
	return nil, nil
}

func (e *Event) UpdateCoverImage(image []byte) (*Photo, *RestError) {
	// TODO
	return nil, nil
}

func (e *Event) ListNewsFeed(page int, size int) ([]Feed, *RestError) {
	// TODO
	return nil, nil
}

func (e *Event) AddNewsFeedMessage(feedPost *FeedPost) (*Feed, *RestError) {
	// TODO
	return nil, nil
}

func (e *Event) Update() *RestError {
	// TODO
	return nil
}

func (e *Event) Cancel() *RestError {
	// TODO
	return nil
}

func (e *Event) ConfirmParticipation() *RestError {
	// TODO
	return nil
}

func (e *Event) Participate() *RestError {
	// TODO
	return nil
}

func (e *Event) CancelParticipation() *RestError {
	// TODO
	return nil
}
