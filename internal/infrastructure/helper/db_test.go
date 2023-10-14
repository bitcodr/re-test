package helper

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitcodr/re-test/internal/infrastructure/config"
)

func TestNewMemory(t *testing.T) {
	// Test case 1: Valid configuration
	cfg := &config.Connection{
		Name: config.MEMORY,
	}

	conn, err := NewMemory(context.Background(), cfg)

	// Assert that the returned error is nil
	assert.Nil(t, err)

	// Assert that the connection map is not nil
	assert.NotNil(t, conn)

	// Add additional assertions based on the expected behavior of the function

	// Test case 2: Empty configuration
	conn, err = NewMemory(context.Background(), nil)

	// Assert that the returned error is not nil
	assert.NotNil(t, err)

	// Assert that the connection map is nil
	assert.Nil(t, conn)

	// Assert that the error message is correct
	assert.EqualError(t, err, "config is empty")

	// Add additional assertions based on the expected behavior of the function
}
