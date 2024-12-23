package log

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "transaction_log"

	idColumn        = "id"
	timestampColumn = "timestamp"
	logColumn       = "log"
)

type repo struct {
	db db.Client
}

// NewRepository creates new object of repository layer.
func NewRepository(db db.Client) repository.LogRepository {
	return &repo{db: db}
}

func (r *repo) Log(ctx context.Context, text *model.Log) error {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, logColumn).
		Values(text.ID, text.Text).
		Suffix("RETURNING " + idColumn)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "log_repository.Log",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
