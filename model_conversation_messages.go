package msa

type ConversationMessages struct {
	Model
	TotalUnreads   int                   `json:"total_unreads"`
	Samples        []ConversationMessage `json:"samples"`
	ConversationId int64                 `json:"conversation_id"`
}

func (c *ConversationMessages) List(page int, size int) ([]ConversationMessage, *RestError) {
	// TODO
	return nil, nil
}

func (c *ConversationMessages) ListAndConsume(page int, size int) ([]ConversationMessage, *RestError) {
	// TODO
	return nil, nil
}
