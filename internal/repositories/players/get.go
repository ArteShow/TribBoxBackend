package players

import (
	"context"
)

func (r *PlayerRepository) GetAll(ctx context.Context) ([]Player, error) {
	query := `SELECT uuid, username, last_seen, created_at FROM players ORDER BY created_at ASC`
	rows, err := r.DB.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var p Player
		if err = rows.Scan(
			&p.UUID,
			&p.Username,
			&p.LastSeen,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}

		players = append(players, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

func (r *PlayerRepository) GetByUUID(ctx context.Context, uuid string) (Player, error) {
	query := `SELECT uuid, username, last_seen, created_at FROM players WHERE uuid = $1`
	row := r.DB.Conn.QueryRow(ctx, query, uuid)

	var player Player
	if err := row.Scan(
		&player.UUID,
		&player.Username,
		&player.LastSeen,
		&player.CreatedAt,
	); err != nil {
		return Player{}, err
	}

	return player, nil
}
