package utils

import (
	"strings"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                                 Normal UUID                                */
/* -------------------------------------------------------------------------- */
func UUID() string {
	id := uuid.New().String()
	return id
}

/* -------------------------------------------------------------------------- */
/*                              UUID with prefix                              */
/* -------------------------------------------------------------------------- */
func UUIDWithPrefix(prefix string) string {
	id := uuid.New().String()
	id = prefix + "_" + id
	id = strings.ReplaceAll(id, "-", "")
	return id
}
