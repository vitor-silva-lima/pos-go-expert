package database

import "gorm.io/gorm"

type ConnectionAdapter interface {
	Connect() (*gorm.DB, error)
}
