package database

import (
	"context"
	_ "embed"
)

//go:embed schema.sql
var dll string

func (q *Queries) Migrate() error {
	ctx := context.Background()

	_, err := q.db.ExecContext(ctx, dll)

	return err
}
