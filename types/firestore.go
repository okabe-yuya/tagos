package types

import (
	"time"
)

// Enum: tagos kinds
type Kind int
const (
	Normal Kind = iota
	Special Kind = iota
)

type TagosNest struct {
	Year string
	Month string
	Day string
	Vote string
}

type TagosRecord struct {
	Sender string
	Receiver string
	Kind Kind
	CreatedAt time.Time
}