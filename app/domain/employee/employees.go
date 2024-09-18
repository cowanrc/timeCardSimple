package employee

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type Employee struct {
	options Options
}

type CreateParams struct {
	FirstName string
	LastName  string
	Email     string
}

// // CreateParamsWithPasswordStrings are the
// // externally-facing parameters that include
// // fields for a user-provided password.
// type CreateParamsWithPasswordStrings struct {
// 	CreateParams
// 	Password        string
// 	ConfirmPassword string
// }

// type CreateParamsWithPasswordHash struct {
// 	CreateParams
// 	PasswordHash []byte
// }

func New(createParams CreateParams) (*Employee, error) {
	employeeID, err := id.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return NewWithOptions(Options{
		ID:        employeeID,
		FirstName: createParams.FirstName,
		LastName:  createParams.LastName,
		Email:     createParams.Email,
		CreatedAt: now,
		UpdatedAt: now,
		// PasswordHash: createParams.PasswordHash,
	})
}

func NewWithOptions(options Options) (*Employee, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	return &Employee{
		options: options,
	}, nil
}
