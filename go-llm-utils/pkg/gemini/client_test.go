package gemini

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewGenAIClientWithAPIKey tests the internal client creation logic.
// By testing the internal function, we can avoid manipulating environment variables,
// which allows the tests to be run in parallel safely.
func TestNewGenAIClientWithAPIKey(t *testing.T) {
	t.Parallel() // Mark the parent test as parallelizable.
	ctx := context.Background()

	t.Run("success when API key is set", func(t *testing.T) {
		t.Parallel() // Mark the subtest as parallelizable.

		client, err := newGenAIClientWithAPIKey(ctx, "fake-api-key")

		require.NoError(t, err)
		require.NotNil(t, client)
	})

	t.Run("error when API key is empty", func(t *testing.T) {
		t.Parallel() // Mark the subtest as parallelizable.

		client, err := newGenAIClientWithAPIKey(ctx, "")

		require.Error(t, err)
		assert.Nil(t, client)
		assert.EqualError(t, err, "API key cannot be empty")
	})
}

// TestNewGenAIClient tests the public constructor that reads from the environment.
// These tests cannot run in parallel because they manipulate the global process state
// (environment variables).
func TestNewGenAIClient(t *testing.T) {
	ctx := context.Background()

	t.Run("success when API key is set via environment variable", func(t *testing.T) {
		// t.Setenv automatically handles cleanup.
		t.Setenv("GEMINI_API_KEY", "fake-key-from-env")

		client, err := NewGenAIClient(ctx)
		require.NoError(t, err)
		require.NotNil(t, client)
	})

	t.Run("error when API key environment variable is not set", func(t *testing.T) {
		const apiKeyName = "GEMINI_API_KEY"
		originalValue, wasSet := os.LookupEnv(apiKeyName)

		// Restore the original value after the test runs.
		if wasSet {
			t.Cleanup(func() {
				assert.NoError(t, os.Setenv(apiKeyName, originalValue))
			})
		}

		// Unset the variable for this specific test case.
		require.NoError(t, os.Unsetenv(apiKeyName))

		client, err := NewGenAIClient(ctx)
		require.Error(t, err)
		assert.Nil(t, client)
		assert.EqualError(t, err, "Environment variable GEMINI_API_KEY not set")
	})

	t.Run("error when API key environment variable is set but empty", func(t *testing.T) {
		t.Setenv("GEMINI_API_KEY", "")

		client, err := NewGenAIClient(ctx)
		require.Error(t, err)
		assert.Nil(t, client)
		assert.EqualError(t, err, "API key cannot be empty")
	})
}
