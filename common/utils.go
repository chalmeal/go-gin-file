package common

import (
	"github.com/google/uuid"
)

func UuId() string {
	id := uuid.New()
	return id.String()
}
