package internal

import (
	"engine/modules/uuid"
	uuidSource "github.com/google/uuid"
)

// impl

type factory struct{}

func NewFactory() uuid.Factory {
	return &factory{}
}

func (factory *factory) NewUUID() uuid.UUID {
	return uuid.UUID(uuidSource.New())
}

func (factory *factory) NewUUIDFromString(seed string) uuid.UUID {
	return uuid.UUID(uuidSource.NewSHA1(uuidSource.NameSpaceURL, []byte(seed)))
}
