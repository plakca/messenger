package domain

import (
	"time"
	"github.com/google/uuid"
)

type Message struct {
	Text string
	CreatedAt time.Time
	AuthorID uuid.UUID
	
	Meta MessageMeta
}

type MessageMeta struct {
	AuthorName string
}



