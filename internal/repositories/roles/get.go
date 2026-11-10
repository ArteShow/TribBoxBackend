package roles

import (
	"context"
)

func (r *RoleRepository) GetAll(ctx context.Context) ([]Role, error) {
	query := `SELECT id, name, description, created_at FROM roles ORDER BY id ASC`
	rows, err := r.DB.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []Role{}
	for rows.Next() {
		var role Role
		if err = rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
		); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepository) GetByID(ctx context.Context, id int64) (Role, error) {
	query := `SELECT id, name, description, created_at FROM roles WHERE id = $1`
	row := r.DB.Conn.QueryRow(ctx, query, id)

	var role Role
	if err := row.Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
	); err != nil {
		return Role{}, err
	}

	return role, nil
}
