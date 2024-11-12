package idtest

import (
	"timeCardSimple/app/domain/id"

	"github.com/google/uuid"
)

func MustNew() id.ID {
	uuid, _ := uuid.NewRandom()
	id := id.ID(uuid)
	return id
}
