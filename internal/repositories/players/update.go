package players

import (
	"context"
	"errors"
)

func (r *PlayerRepository) UpdatePlayer(ctx context.Context, uuid, username string) error {
	sql := `UPDATE players SET username = $1 WHERE uuid = $2`
	tag, err := r.DB.Conn.Exec(ctx, sql, username, uuid)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("player not found")
	}

	return nil
}
