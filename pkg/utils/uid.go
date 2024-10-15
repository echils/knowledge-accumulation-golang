package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/bwmarrin/snowflake"
)

func Random32UUID() string {
	uid := uuid.NewString()
	return strings.ReplaceAll(uid, "-", "")
}

func SnowflakeID() string {
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Printf("Snowflake node load failed: %v", err)
		os.Exit(1)
	}
	return n.Generate().String()
}
