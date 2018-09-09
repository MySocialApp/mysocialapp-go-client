package msa

func newClientService(session *Session) *ClientService {
	return &ClientService{Session: session}
}

type ClientService struct {
	Session *Session

	account       *RestAccount
	conversation  *RestConversation
	dynamicEntity *RestDynamicEntity
	event         *RestEvent
	feed          *RestFeed
	friend        *RestFriend
	group         *RestGroup
	login         *RestLogin
	notification  *RestNotification
	photo         *RestPhoto
	register      *RestRegister
	reset         *RestReset
	search        *RestSearch
	status        *RestStatus
	user          *RestUser
}

func (c *ClientService) Account() *RestAccount {
	if c.account == nil {
		c.account = &RestAccount{Configuration: c.Session.Configuration}
	}
	return c.account
}

func (c *ClientService) Conversation() *RestConversation {
	if c.conversation == nil {
		c.conversation = &RestConversation{Configuration: c.Session.Configuration}
	}
	return c.conversation
}

func (c *ClientService) DynamicEntity() *RestDynamicEntity {
	if c.dynamicEntity == nil {
		c.dynamicEntity = &RestDynamicEntity{Configuration: c.Session.Configuration}
	}
	return c.dynamicEntity
}

func (c *ClientService) Event() *RestEvent {
	if c.event == nil {
		c.event = &RestEvent{Configuration: c.Session.Configuration}
	}
	return c.event
}

func (c *ClientService) Feed() *RestFeed {
	if c.feed == nil {
		c.feed = &RestFeed{Configuration: c.Session.Configuration}
	}
	return c.feed
}

func (c *ClientService) Friend() *RestFriend {
	if c.friend == nil {
		c.friend = &RestFriend{Configuration: c.Session.Configuration}
	}
	return c.friend
}

func (c *ClientService) Group() *RestGroup {
	if c.group == nil {
		c.group = &RestGroup{Configuration: c.Session.Configuration}
	}
	return c.group
}

func (c *ClientService) Login() *RestLogin {
	if c.login == nil {
		c.login = &RestLogin{Configuration: c.Session.Configuration}
	}
	return c.login
}

func (c *ClientService) Notification() *RestNotification {
	if c.notification == nil {
		c.notification = &RestNotification{Configuration: c.Session.Configuration}
	}
	return c.notification
}

func (c *ClientService) Photo() *RestPhoto {
	if c.photo == nil {
		c.photo = &RestPhoto{Configuration: c.Session.Configuration}
	}
	return c.photo
}

func (c *ClientService) Register() *RestRegister {
	if c.register == nil {
		c.register = &RestRegister{Configuration: c.Session.Configuration}
	}
	return c.register
}

func (c *ClientService) Reset() *RestReset {
	if c.reset == nil {
		c.reset = &RestReset{Configuration: c.Session.Configuration}
	}
	return c.reset
}

func (c *ClientService) Search() *RestSearch {
	if c.search == nil {
		c.search = &RestSearch{Configuration: c.Session.Configuration}
	}
	return c.search
}

func (c *ClientService) Status() *RestStatus {
	if c.status == nil {
		c.status = &RestStatus{Configuration: c.Session.Configuration}
	}
	return c.status
}

func (c *ClientService) User() *RestUser {
	if c.user == nil {
		c.user = &RestUser{Configuration: c.Session.Configuration}
	}
	return c.user
}
