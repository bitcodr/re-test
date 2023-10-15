package helper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitcodr/re-test/internal/infrastructure/config"
	"github.com/bitcodr/re-test/internal/infrastructure/helper"
)

func TestNewMemory(t *testing.T) {
	// Test case 1: Valid configuration
	cfg := &config.Connection{
		Name: config.MEMORY,
	}

	_, err := helper.NewMemory(context.Background(), cfg)

	// Assert that the returned error is nil
	assert.Nil(t, err)

	// Add additional assertions based on the expected behavior of the function

	// Test case 2: Empty configuration
	conn, err := helper.NewMemory(context.Background(), nil)

	// Assert that the returned error is not nil
	assert.NotNil(t, err)

	// Assert that the connection map is nil
	assert.Nil(t, conn)

	// Assert that the error message is correct
	assert.EqualError(t, err, "config is empty")
}
