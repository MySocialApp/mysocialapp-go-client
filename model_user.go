package msa

import "time"

type User struct {
	Base                   `json:"-"`
	Id             *int64  `json:"id"`
	IdStr          *string `json:"id_str"`
	FirstName      string  `json:"first_name"`
	LastName       *string `json:"last_name"`
	UpdatedDate    *string `json:"updated_date"`
	Email          string  `json:"email"`
	ValidatedEmail *bool   `json:"validated_email"`
	Gender         *Gender `json:"gender"`
	DateOfBirth    *string `json:"date_of_birth"`
	Presentation   *string `json:"presentation"`

	ProfilePhoto        *Photo        `json:"profile_photo"`
	ProfileCoverPhoto   *Photo        `json:"profile_cover_photo"`
	LivingLocation      *Location     `json:"living_location"`
	CurrentStatus       *Status       `json:"current_status"`
	Authorities         []string      `json:"authorities"`
	AccountEnabled      *bool         `json:"account_enabled"`
	AccountExpired      *bool         `json:"account_expired"`
	Flag                *Flag         `json:"flag"`
	UserSettings        *UserSettings `json:"user_settings"`
	UserStat            *UserStat     `json:"user_stat"`
	IsFriend            *bool         `json:"is_friend"`
	IsRequestedAsFriend *bool         `json:"is_requested_as_friend"`
	ExternalId          *string       `json:"external_id"`
	CustomFields        []CustomField `json:"custom_fields"`
}

func (u *User) RemoveFriend() *RestError {
	return nil
}

func (u *User) RequestAsFriend() (*User, *RestError) {
	return nil, nil
}

func (u *User) CancelFriendRequest() *RestError {
	return nil
}

func (u *User) AcceptFriendRequest() (*User, *RestError) {
	return nil, nil
}

func (u *User) RefuseFriendRequest() *RestError {
	return nil
}

func (u *User) ListFriend(page int, size int) ([]User, *RestError) {
	return nil, nil
}

func (u *User) ListNewsFeed(page int, size int) ([]Feed, *RestError) {
	return nil, nil
}

func (e *User) AddNewsFeedMessage(feedPost *FeedPost) (*Feed, *RestError) {
	// TODO
	return nil, nil
}

func (u *User) SendPrivateMessage(message *ConversationMessagePost) (*ConversationMessage, *RestError) {
	return nil, nil
}

func (u *User) ListGroup(page int, size int) ([]Group, *RestError) {
	return nil, nil
}

func (u *User) ListEvent(page int, size int, from *time.Time) ([]Event, *RestError) {
	return nil, nil
}

func (u *User) ListPhotoAlbum(page int, size int) ([]PhotoAlbum, *RestError) {
	return nil, nil
}