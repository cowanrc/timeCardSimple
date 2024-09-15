package employee

import (
	"time"

	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/lib/clone"
)

type Options struct {
	ID           id.ID
	FirstName    string
	LastName     string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PasswordHash []byte
}

func (o Options) validate() error {
	if o.ID == (id.ID{}) {
		return ErrUnsetID
	}
	if o.FirstName == "" {
		return ErrUnsetFirstName
	}
	if o.LastName == "" {
		return ErrUnsetLastName
	}

	if o.Email == "" {
		return ErrUnsetEmail
	}

	if o.CreatedAt.IsZero() {
		return ErrUnsetCreatedAt
	}

	if o.UpdatedAt.IsZero() {
		return ErrUnsetUpdatedAt
	}
	return nil
}

func (o Options) DeepClone() Options {
	clonedOptions := o

	clonedOptions.PasswordHash = clone.Slice(o.PasswordHash)

	return clonedOptions
}
