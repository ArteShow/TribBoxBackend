package roles

import "context"

func (r *RoleRepository) CreateRole(ctx context.Context, name, description string) (int64, error) {
	sql := `INSERT INTO roles (name, description)
			VALUES ($1, $2)
			RETURNING id
			ON CONFLICT (name) DO NOTHING;`

	var id int64
	err := r.DB.Conn.QueryRow(ctx, sql, name, description).Scan(&id)
	return id, err
}
