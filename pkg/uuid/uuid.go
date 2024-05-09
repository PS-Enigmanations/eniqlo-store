package uuid

import (
	"github.com/google/uuid"
)

func New() string {
	uuid := uuid.NewString()
	return uuid
}
