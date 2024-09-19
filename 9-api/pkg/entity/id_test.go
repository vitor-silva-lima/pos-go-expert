package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotNil(t, id)
}

func TestStringToID(t *testing.T) {
	id, err := StringToID("550e8400-e29b-41d4-a716-446655440000")
	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestStringToIDWhenInvalid(t *testing.T) {
	id, err := StringToID("invalid")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid UUID length: 7")
	assert.Equal(t, id, ID{})
}
