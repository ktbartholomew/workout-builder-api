package util

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
)

// ToJSON converts its argument to JSON and returns it as a string
func ToJSON(v interface{}) string {
	marshalled, err := json.Marshal(v)

	if err != nil {
		return "{}"
	}

	return string(marshalled)
}

// NewUUID exports a stringified UUIDv4
func NewUUID() string {
	u := make([]byte, 16)

	rand.Read(u)

	// Set the version byte to 4
	u[6] = (u[6] & 0xf) | 4<<4

	// Set the RFC4122 variant character to 8..b
	// This means changing the first four bits of u[8] and keeping the second four
	u[8] = (u[8] & 0xF) | ((u[8]&0x3 | 0x8) << 4)

	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
