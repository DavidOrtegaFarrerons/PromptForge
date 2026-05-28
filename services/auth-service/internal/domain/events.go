package domain

import "time"

type UserRegisteredEvent struct {
	UserID     string    `json:"user_id"`
	Email      string    `json:"email"`
	OccurredAt time.Time `json:"occurred_at"`
}

func NewUserRegisteredEvent(userID UserID, email Email, occurredAt time.Time) UserRegisteredEvent {
	return UserRegisteredEvent{
		UserID:     string(userID),
		Email:      email.Value(),
		OccurredAt: occurredAt,
	}
}
