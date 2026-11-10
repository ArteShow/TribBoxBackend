package roles

import (
	"context"
	"errors"
)

func (r *RoleRepository) UpdateRole(ctx context.Context, name, description string, id int64) error {
	query := `UPDATE roles SET name = $1, description = $2 WHERE id = $3;`
	tag, err := r.DB.Conn.Exec(ctx, query, name, description, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("role not found")
	}

	return nil
}
