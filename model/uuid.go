package model

import (
	"strings"

	"github.com/google/uuid"
)

// 生成UUID
func RandomUUID() (id string) {
	s := uuid.NewString()
	return strings.ReplaceAll(s, "-", "")
}
