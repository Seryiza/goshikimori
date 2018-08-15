package models

import (
	"time"
)

const (
	// MessageKindPrivate is one of InnerMessage.Kind
	MessageKindPrivate = "Private"

	// Message types
	MessageTypeNews          = "news"
	MessageTypeNotifications = "notifications"
)

// Message of Shikimori
type Message struct {
	Frontend bool         `json:"frontend"`
	Inner    InnerMessage `json:"message"`
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

// MarkReadMessages - Mark messages as read or unread.
// https://shikimori.org/api/doc/1.0/messages/read_all
type MarkReadMessages struct {
	// IDs list separated by comma.
	// Ex., "17,18,987654".
	IDs string `json:"ids"`

	// Ex., "1"
	IsRead string `json:"is_read"`
}

// ReadAllMessages - Mark all messages as read
// https://shikimori.org/api/doc/1.0/messages/read_all
type ReadAllMessages struct {
	Frontend     bool   `json:"frontend"`
	MessagesType string `json:"type"`
}

// DeleteAllMessages - Delete all messages
// https://shikimori.org/api/doc/1.0/messages/delete_all
type DeleteAllMessages struct {
	Frontend     bool   `json:"frontend"`
	MessagesType string `json:"type"`
}
