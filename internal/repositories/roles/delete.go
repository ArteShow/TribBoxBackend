package roles

import (
	"context"
	"errors"
)

func (r *RoleRepository) DeleteRole(ctx context.Context, id int64) error {
	query := `DELETE FROM roles WHERE id = $1`
	tag, err := r.DB.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("role not found")
	}

	return err
}
