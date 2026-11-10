package roles

import "context"

func (r *RoleRepository) CreateRole(ctx context.Context, name, description string) error {
	sql := `INSERT INTO roles (name, description)
			VALUES ($1, $2)
			ON CONFLICT (name) DO NOTHING;`

	_, err := r.DB.Conn.Exec(ctx, sql, name, description)
	return err
}
