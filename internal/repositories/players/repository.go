package players

import (
	"time"

	"github.com/ArteShow/TribBoxBackend/internal/database"
)

type PlayerRepository struct {
	DB *database.Database
}

func NewPlayerRepository(db *database.Database) *PlayerRepository {
	return &PlayerRepository{
		DB: db,
	}
}

type Player struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	LastSeen  time.Time `json:"last_seen"`
	CreatedAt time.Time `json:"created_at"`
}
