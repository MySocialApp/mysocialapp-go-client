package msa

type Conversation struct {
	Members  []User               `json:"members"`
	Messages *ConversationMessage `json:"messages"`
}

func (c *Conversation) SendMessage(message ConversationMessagePost) (*ConversationMessage, *RestError) {
	// TODO
	return nil, nil
}

func (c *Conversation) KickMember(user *User) *RestError {
	// TODO
	return nil
}

func (c *Conversation) AddMember(user *User) (*User, *RestError) {
	// TODO
	return nil, nil
}

func (c *Conversation) Delete() *RestError {
	// TODO
	return nil
}

func (c *Conversation) Quit() *RestError {
	return c.Delete()
}

func (c *Conversation) Update() *RestError {
	// TODO
	return nil
}