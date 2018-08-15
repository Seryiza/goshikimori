package models

import (
	"time"
)

const (
	// MessageKindPrivate is one of InnerMessage.Kind
	MessageKindPrivate = "Private"
)

// Message of Shikimori
type Message struct {
	Inner InnerMessage `json:"message"`
}

// InnerMessage of Message struct
type InnerMessage struct {
	Body string `json:"body"`
	Kind string `json:"kind"`

	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
}

// MessageResult of POST-request Message struct
type MessageResult struct {
	ID   int32  `json:"id"`
	Kind string `json:"kind"`
	Read bool   `json:"read"`

	Body     string `json:"body"`
	HTMLBody string `json:"html_body"`

	CreatedAt time.Time `json:"created_at"`

	// todo: посмотреть, что это.
	Linked interface{}

	From User `json:"from"`
	To   User `json:"to"`
}
