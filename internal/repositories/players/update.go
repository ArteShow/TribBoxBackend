package players

import (
	"context"
	"errors"
)

func (r *PlayerRepository) UpdatePlayerUsername(ctx context.Context, uuid, username string) error {
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

func (r *PlayerRepository) TouchPlayer(ctx context.Context, uuid string) error {
	sql := `UPDATE players SET last_seen = NOW() WHERE uuid = $1`
	tag, err := r.DB.Conn.Exec(ctx, sql, uuid)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("player not found")
	}

	return nil
}
