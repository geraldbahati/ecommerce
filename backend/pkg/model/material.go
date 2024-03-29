package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Material struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	CreatedAt   time.Time    `json:"created_at"`
	LastUpdated sql.NullTime `json:"last_updated"`
}
