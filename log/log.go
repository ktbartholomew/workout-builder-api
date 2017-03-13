package log

import (
	"os"
	"time"

	"github.com/ktbartholomew/workout-builder-api/util"
)

// Message defines the
type Message struct {
	Name      string
	Timestamp time.Time
	Event     string
	LogLevel  string
	Data      string
	RequestID string
}

// NewMessage creates a Message with default values
func NewMessage(m Message) Message {
	message := Message{
		Name:      "workout-builder-api",
		Timestamp: time.Now(),
		LogLevel:  "info"}

	if m.Event != "" {
		message.Event = m.Event
	}

	if m.LogLevel != "" {
		message.LogLevel = m.LogLevel
	}

	if m.RequestID != "" {
		message.RequestID = m.RequestID
	}

	return message
}

// Log makes it easier to log things
func Log(m Message) {
	os.Stdout.WriteString(util.ToJSON(NewMessage(m)) + "\n")
}
