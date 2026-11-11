package players

import (
	"context"
	"errors"
)

func (r *PlayerRepository) DeletePlayer(ctx context.Context, uuid string) error {
	query := `DELETE FROM players WHERE uuid = $1`
	tag, err := r.DB.Conn.Exec(ctx, query, uuid)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("player not found")
	}

	return nil
}
