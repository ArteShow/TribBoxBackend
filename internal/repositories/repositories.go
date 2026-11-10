package repositories

import (
	"github.com/ArteShow/TribBoxBackend/internal/database"
	"github.com/ArteShow/TribBoxBackend/internal/repositories/roles"
)

type Repositories struct {
	RoleRepository *roles.RoleRepository
}

func NewRepositories(db *database.Database) *Repositories {
	return &Repositories{
		RoleRepository: roles.NewRoleRepository(db),
	}
}
