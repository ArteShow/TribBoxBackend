package roles

import (
	"time"

	"github.com/ArteShow/TribBoxBackend/internal/database"
)

type RoleRepository struct {
	DB *database.Database
}

func NewRoleRepository(db *database.Database) *RoleRepository {
	return &RoleRepository{
		DB: db,
	}
}

type Role struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
