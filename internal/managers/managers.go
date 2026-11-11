package managers

import (
	"github.com/ArteShow/TribBoxBackend/internal/managers/players"
	"github.com/ArteShow/TribBoxBackend/internal/managers/roles"
	"github.com/ArteShow/TribBoxBackend/internal/repositories"
)

type Managers struct {
	RoleManager   *roles.RoleManager
	PlayerManager *players.PlayerManager
}

func NewManagerHub(repositories *repositories.Repositories) *Managers {
	return &Managers{
		RoleManager:   roles.NewRoleManager(repositories.RoleRepository),
		PlayerManager: players.NewPlayerManager(repositories.PlayerRepository),
	}
}
