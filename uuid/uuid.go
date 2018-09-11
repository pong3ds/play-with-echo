package uuid

import (
	guuid "github.com/google/uuid"
)

// IUUID create the uuid
type IUUID interface {
	GetUUID() string
}

// UUID is the utility to create udid
type UUID struct {
}

// NewUUID create the UUID
func NewUUID() *UUID {
	return &UUID{}
}

// GetUUID return uuid string
func (*UUID) GetUUID() string {
	id, _ := guuid.NewUUID()
	return id.String()
}
