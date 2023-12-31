package helper

import (
	"context"
	"fmt"

	"github.com/bitcodr/re-test/internal/infrastructure/config"
)

// NewMemory start connection for memory - just an example how data can be persisted
func NewMemory(_ context.Context, cfg *config.Connection) ([]int, error) {
	if cfg == nil {
		return nil, fmt.Errorf("%s", "config is empty")
	}

	// just an example how data can be persisted, it can be any connection
	var conn []int

	return conn, nil
}
