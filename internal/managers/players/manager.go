package players

import (
	"github.com/ArteShow/TribBoxBackend/internal/repositories/players"
)

type PlayerManager struct {
	Repository *players.PlayerRepository
}

func NewPlayerManager(repo *players.PlayerRepository) *PlayerManager {
	return &PlayerManager{
		Repository: repo,
	}
}
