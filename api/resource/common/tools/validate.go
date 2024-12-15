package tools

import (
	"github.com/google/uuid"
)

func IsValidId(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func RandomId() string {
	return uuid.New().String()
}