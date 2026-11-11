package players

import "context"

func (r *PlayerRepository) CreatePlayer(ctx context.Context, uuid, username string) error {
	sql := `INSERT INTO players
			VALUES (uuid, username)
			ON CONFLICT (name) DO NOTHING;`

	_, err := r.DB.Conn.Exec(ctx, sql, uuid, username)
	return err
}
