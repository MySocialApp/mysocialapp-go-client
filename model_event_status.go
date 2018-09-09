package msa

type EventStatus string

const (
	EventStatusWantToParticipate                EventStatus = "WANT_TO_PARTICIPATE"
	EventStatusWaitingConfirmation              EventStatus = "WAITING_CONFIRMATION"
	EventStatusConfirmed                        EventStatus = "CONFIRMED"
	EventStatusWaitingForFreeSeat               EventStatus = "WAITING_FOR_FREE_SEAT"
	EventStatusNoResponse                       EventStatus = "NO_RESPONSE"
	EventStatusNotAvailable                     EventStatus = "NOT_AVAILABLE"
	EventStatusHasCancelled                     EventStatus = "HAS_CANCELLED"
	EventStatusHasCancelledAfterHavingConfirmed EventStatus = "HAS_CANCELLED_AFTER_HAVING_CONFIRMED"
)
