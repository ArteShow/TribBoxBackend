package managers

import (
	"github.com/ArteShow/TribBoxBackend/internal/managers/roles"
	"github.com/ArteShow/TribBoxBackend/internal/repositories"
)

type Managers struct {
	RoleManager *roles.RoleManager
}

func NewManagerHub(repositories *repositories.Repositories) *Managers {
	return &Managers{
		RoleManager: roles.NewRoleManager(repositories.RoleRepository),
	}
}
