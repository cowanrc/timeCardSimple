package employee

import (
	"reflect"
	"testing"
	"time"
	"timeCardSimple/app/domain/id/idtest"
)

func Test_Options_validate(t *testing.T) {
	options := Options{}

	if err := options.validate(); err != ErrUnsetID {
		t.Errorf("incorrect error, received: %v, expected: %v", err, ErrUnsetID)
	}

	options.ID = idtest.MustNew()

	if err := options.validate(); err != ErrUnsetFirstName {
		t.Errorf("incorrect error, received: %v, expected: %v", err, ErrUnsetID)
	}

	options.FirstName = "first"

	if err := options.validate(); err != ErrUnsetLastName {
		t.Errorf("incorrect error, received: %v, expected: %v", err, ErrUnsetID)
	}

	options.LastName = "last"

	if err := options.validate(); err != ErrUnsetEmail {
		t.Errorf("incorrect error, received: %v, expected: %v", err, ErrUnsetID)
	}

	options.Email = "test@email.com"

	now := time.Now()

	if err := options.validate(); err != ErrUnsetCreatedAt {
		t.Errorf("incorrect error, received: %v, expected: %v", err, ErrUnsetID)
	}

	options.CreatedAt = now

	if err := options.validate(); err != ErrUnsetUpdatedAt {
		t.Errorf("incorrect error, received: %v, expected: %v", err, ErrUnsetID)
	}

	options.UpdatedAt = now
}

func Test_Options_deepClone(t *testing.T) {
	originalOptions := Options{
		ID:        idtest.MustNew(),
		FirstName: "first",
		LastName:  "last",
		Email:     "test@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	clonedOptions := originalOptions.DeepClone()

	if reflect.DeepEqual(originalOptions, clonedOptions) {
		t.Errorf("incorrect deep clone, received: %+v, expected: %+v", clonedOptions, originalOptions)
	}
}
