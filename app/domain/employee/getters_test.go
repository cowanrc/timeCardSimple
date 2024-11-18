package employee

import (
	"testing"
	"time"

	"timeCardSimple/app/domain/id/idtest"
)

func TestEmployee_AllGettersReturnTheCorrectValue(t *testing.T) {
	id := idtest.MustNew()
	email := "test@email.com"
	firstName := "first"
	lastName := "last"
	now := time.Now()

	u := &Employee{
		options: Options{
			ID:        id,
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	if !u.ID().Equal(id) {
		t.Fatalf("incorrect id, received: %v, expected: %v", u.ID(), id)
	}

	if u.Email() != email {
		t.Fatalf("incorrect email, received: %s, expected: %s", u.Email(), email)
	}

	if u.FirstName() != firstName {
		t.Fatalf("incorrect firstName, received: %s, expected: %s", u.FirstName(), firstName)
	}

	if u.LastName() != lastName {
		t.Fatalf("incorrect lastName, received: %s, expected: %s", u.LastName(), lastName)
	}

	if !u.CreatedAt().Equal(now) {
		t.Fatalf("incorrect createdAt, received: %v, expected: %v", u.CreatedAt(), now)
	}

	if !u.UpdatedAt().Equal(now) {
		t.Fatalf("incorrect updatedAt, received: %v, expected: %v", u.UpdatedAt(), now)
	}
}
