package roles

import (
	"context"
	"errors"

	"github.com/ArteShow/TribBoxBackend/internal/repositories/roles"
)

var (
	ErrInvalidRole   = errors.New("invalid role")
	ErrProtectedRole = errors.New("cannot modify system role")
)

type RoleManager struct {
	Repository *roles.RoleRepository
}

func NewRoleManager(repo *roles.RoleRepository) *RoleManager {
	return &RoleManager{
		Repository: repo,
	}
}

func (m *RoleManager) CreateRole(ctx context.Context, name, description string) (int64, error) {
	if name == "" || description == "" {
		return 0, ErrInvalidRole
	}

	id, err := m.Repository.CreateRole(ctx, name, description)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *RoleManager) DeleteRole(ctx context.Context, id int64) error {
	role, err := m.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if role.Name == "ADMIN" || role.Name == "MODERATOR" || role.Name == "PLAYER" {
		return ErrProtectedRole
	}

	return m.Repository.DeleteRole(ctx, id)
}

func (m *RoleManager) GetAllRoles(ctx context.Context) ([]roles.Role, error) {
	return m.Repository.GetAll(ctx)
}

func (m *RoleManager) GetRoleByID(ctx context.Context, id int64) (roles.Role, error) {
	return m.Repository.GetByID(ctx, id)
}

func (m *RoleManager) UpdateRole(ctx context.Context, name, description string, id int64) error {
	role, err := m.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if role.Name == "ADMIN" || role.Name == "MODERATOR" || role.Name == "PLAYER" {
		return ErrProtectedRole
	}

	if name == "" || description == "" {
		return ErrInvalidRole
	}

	return m.Repository.UpdateRole(ctx, name, description, id)
}
