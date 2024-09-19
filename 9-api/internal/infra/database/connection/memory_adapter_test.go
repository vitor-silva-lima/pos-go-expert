package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryConnectionAdapter_Connect(t *testing.T) {
	adapter := NewMemoryConnectionAdapter()
	db, err := adapter.Connect()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
