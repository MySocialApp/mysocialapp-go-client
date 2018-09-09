package msa

func newSession(conf *Configuration) *Session {
	session := Session{
		Configuration: conf,
	}
	session.initClient()
	return &session
}

type Session struct {
	Configuration *Configuration
	ClientService *ClientService

	account      *FluentAccount
	conversation *FluentConversation
	event        *FluentEvent
	friend       *FluentFriend
	group        *FluentGroup
	newsFeed     *FluentNewsFeed
	notification *FluentNotification
	photo        *FluentPhoto
	photoAlbum   *FluentPhotoAlbum
	user         *FluentUser
}

func (s *Session) initClient() {
	s.ClientService = newClientService(s)
}

func (s *Session) Account() *FluentAccount {
	if s.account == nil {
		s.account = &FluentAccount{Session: s}
	}
	return s.account
}

func (s *Session) Conversation() *FluentConversation {
	if s.conversation == nil {
		s.conversation = &FluentConversation{Session: s}
	}
	return s.conversation
}

func (s *Session) Event() *FluentEvent {
	if s.event == nil {
		s.event = &FluentEvent{Session: s}
	}
	return s.event
}

func (s *Session) Friend() *FluentFriend {
	if s.friend == nil {
		s.friend = &FluentFriend{Session: s}
	}
	return s.friend
}

func (s *Session) Group() *FluentGroup {
	if s.group == nil {
		s.group = &FluentGroup{Session: s}
	}
	return s.group
}

func (s *Session) NewsFeed() *FluentNewsFeed {
	if s.newsFeed == nil {
		s.newsFeed = &FluentNewsFeed{Session: s}
	}
	return s.newsFeed
}

func (s *Session) Notification() *FluentNotification {
	if s.notification == nil {
		s.notification = &FluentNotification{Session: s}
	}
	return s.notification
}

func (s *Session) Photo() *FluentPhoto {
	if s.photo == nil {
		s.photo = &FluentPhoto{Session: s}
	}
	return s.photo
}

func (s *Session) PhotoAlbum() *FluentPhotoAlbum {
	if s.photoAlbum == nil {
		s.photoAlbum = &FluentPhotoAlbum{Session: s}
	}
	return s.photoAlbum
}
