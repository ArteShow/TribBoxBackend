package migrate

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/ArteShow/TribBoxBackend/internal/database"
	"github.com/jackc/pgx/v5"
)

type Runner struct {
	db *database.Database
}

func NewRunner(db *database.Database) *Runner {
	return &Runner{
		db: db,
	}
}

func (r *Runner) recordMigration(ctx context.Context, tx pgx.Tx, sql string, migration Migration) error {
	_, err := tx.Exec(ctx, sql, migration.Version, migration.Name)
	return err
}

func (r *Runner) EnsureMigrationsTable(ctx context.Context) error {
	sql := `CREATE TABLE IF NOT EXISTS schema_migrations (
    			version BIGINT PRIMARY KEY,
    			name TEXT NOT NULL,
    			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
			);`

	_, err := r.db.Conn.Exec(ctx, sql)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) GetAppliedMigrations(ctx context.Context) ([]Migration, error) {
	query := `SELECT name, version FROM schema_migrations ORDER BY version ASC`
	rows, err := r.db.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appliedMigrations := []Migration{}
	for rows.Next() {
		var migration Migration
		if err := rows.Scan(
			&migration.Name,
			&migration.Version,
		); err != nil {
			return nil, err
		}

		appliedMigrations = append(appliedMigrations, migration)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return appliedMigrations, nil
}

func (r *Runner) LoadMigrations() ([]Migration, error) {
	files, err := os.ReadDir("./migrations")
	if err != nil {
		return nil, err
	}

	migrations := []Migration{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if strings.HasSuffix(fileName, ".sql") {
			if len(fileName) < 9 || fileName[3] != '_' {
				return nil, fmt.Errorf("invalid migration filename: %s", fileName)
			}

			version, err := strconv.ParseInt(fileName[:3], 10, 64)
			if err != nil {
				return nil, err
			}

			name := fileName[4 : len(fileName)-4]
			migrations = append(migrations, Migration{
				Name:    name,
				Version: version,
				Path:    filepath.Join("./migrations", fileName),
			})
		}
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func (r *Runner) ApplyMigration(ctx context.Context, migration Migration) error {
	if migration.Version == 0 {
		return errors.New("invalid migration version")
	}

	body, err := os.ReadFile(migration.Path)
	if err != nil {
		return err
	}

	sql := string(body)

	tx, err := r.db.Conn.Begin(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, sql); err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			return errors.New(err.Error() + " " + txErr.Error())
		}

		return err
	}

	if err = r.recordMigration(ctx,
		tx,
		`INSERT INTO schema_migrations (version, name) VALUES ($1, $2);`,
		migration); err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			return errors.New(err.Error() + " " + txErr.Error())
		}

		return err
	}

	if err = tx.Commit(ctx); err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			return errors.New(err.Error() + " " + txErr.Error())
		}

		return err
	}

	return nil
}
