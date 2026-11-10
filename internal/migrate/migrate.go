package migrate

import (
	"context"
	"slices"

	"github.com/ArteShow/TribBoxBackend/internal/database"
)

func RunMigrations(ctx context.Context, db *database.Database) error {
	runner := NewRunner(db)

	if err := runner.EnsureMigrationsTable(ctx); err != nil {
		return err
	}

	migrations, err := runner.LoadMigrations()
	if err != nil {
		return err
	}

	appliedMigrations, err := runner.GetAppliedMigrations(ctx)
	if err != nil {
		return err
	}

	for _, m := range migrations {
		if !slices.ContainsFunc(appliedMigrations, func(applied Migration) bool {
			return applied.Version == m.Version
		}) {
			if err = runner.ApplyMigration(ctx, m); err != nil {
				return err
			}
		}
	}

	return nil
}
