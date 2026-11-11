package players

import (
	"context"
	"errors"

	"github.com/ArteShow/TribBoxBackend/internal/repositories/players"
	uuid_lib "github.com/google/uuid"
)

var (
	ErrInvalidPlayer = errors.New("invalid player")
)

type PlayerManager struct {
	Repository *players.PlayerRepository
}

func NewPlayerManager(repo *players.PlayerRepository) *PlayerManager {
	return &PlayerManager{
		Repository: repo,
	}
}

func (m *PlayerManager) CreatePlayer(ctx context.Context, uuid, username string) error {
	if username == "" || uuid == "" {
		return ErrInvalidPlayer
	}

	if err := uuid_lib.Validate(uuid); err != nil {
		return ErrInvalidPlayer
	}

	if err := m.Repository.CreatePlayer(ctx, uuid, username); err != nil {
		return err
	}

	return nil
}

func (m *PlayerManager) DeletePlayer(ctx context.Context, uuid string) error {
	if err := uuid_lib.Validate(uuid); err != nil {
		return ErrInvalidPlayer
	}

	return m.Repository.DeletePlayer(ctx, uuid)
}

func (m *PlayerManager) GetAllPlayers(ctx context.Context) ([]players.Player, error) {
	return m.Repository.GetAll(ctx)
}

func (m *PlayerManager) GetPlayerByUUID(ctx context.Context, uuid string) (players.Player, error) {
	if err := uuid_lib.Validate(uuid); err != nil {
		return players.Player{}, ErrInvalidPlayer
	}

	return m.Repository.GetByUUID(ctx, uuid)
}

func (m *PlayerManager) UpdatePlayerUsername(ctx context.Context, uuid, newUsername string) error {
	if err := uuid_lib.Validate(uuid); err != nil {
		return ErrInvalidPlayer
	}

	if newUsername == "" {
		return ErrInvalidPlayer
	}

	return m.Repository.UpdatePlayerUsername(ctx, uuid, newUsername)
}

func (m *PlayerManager) TouchPlayer(ctx context.Context, uuid string) error {
	if err := uuid_lib.Validate(uuid); err != nil {
		return ErrInvalidPlayer
	}

	return m.Repository.TouchPlayer(ctx, uuid)
}
