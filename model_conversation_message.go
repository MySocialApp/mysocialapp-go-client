package msa

type ConversationMessage struct {
	Base
	Message      *string       `json:"message"`
	Photo        *Photo        `json:"photo"`
	Conversation *Conversation `json:"conversation"`
}
