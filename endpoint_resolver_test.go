package veneur

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefinitelyDontResolve(t *testing.T) {
	er := NewEndpointResolver(1.0)

	expectedValue := "1234"
	er.lastResolvedURL = expectedValue

	known, err := er.ResolveEndpoint("testing")
	assert.NoError(t, err, "Error resolving endpoint")

	assert.Equal(t, expectedValue, known, "Failed to resolve address to known value")
}

func TestDefinitelyResolve(t *testing.T) {
	er := NewEndpointResolver(0.0)

	expectedValue := "1234"
	er.lastResolvedURL = expectedValue

	known, err := er.ResolveEndpoint("http://www.example.com:80/")
	assert.NoError(t, err, "Error resolving endpoint")

	assert.NotEqual(t, expectedValue, known, "Failed to resolve address to known value")
}

func TestMaybeResolve(t *testing.T) {
	// Make our random numbers probabilistic
	rand.Seed(1)

	// 0.6046603
	// 0.9405091
	// 0.6645601

	expectedValue := "1234"
	er := NewEndpointResolver(0.7)
	er.lastResolvedURL = expectedValue

	known, err := er.ResolveEndpoint("http://localhost:80/")
	assert.NoError(t, err, "Error resolving endpoint")
	assert.Equal(t, expectedValue, known)

	newVal, err := er.ResolveEndpoint("http://localhost:80/")
	assert.NoError(t, err, "Error resolving endpoint")
	assert.NotEqual(t, expectedValue, newVal)

	assert.True(t, true, true)
}
