package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MemoryConnectionAdapter struct {
}

func NewMemoryConnectionAdapter() *MemoryConnectionAdapter {
	return &MemoryConnectionAdapter{}
}

func (m *MemoryConnectionAdapter) Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
}
