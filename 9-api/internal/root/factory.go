package root

import (
	database "first-api/internal/infra/database/connection"
)

func NewDatabaseConnectionAdapter() database.ConnectionAdapter {
	return database.NewMemoryConnectionAdapter()
}
