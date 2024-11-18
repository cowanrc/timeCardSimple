package employee

import (
	"time"
	"timeCardSimple/app/domain/id"
)

// ID returns the employee ID
func (e *Employee) ID() id.ID {
	return e.options.ID
}

func (e *Employee) FirstName() string {
	return e.options.FirstName
}

func (e *Employee) LastName() string {
	return e.options.LastName
}

func (e *Employee) Email() string {
	return e.options.Email
}

func (e *Employee) CreatedAt() time.Time {
	return e.options.CreatedAt
}

// UpdatedAt returns the User updatedAt field.
func (e *Employee) UpdatedAt() time.Time {
	return e.options.UpdatedAt
}

// PasswordHash returns the User password field.
// func (e *Employee) PasswordHash() []byte {
// 	return clone.Slice(e.options.PasswordHash)
// }
