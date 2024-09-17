package access

import (
	"context"
	"errors"

	accessService "github.com/8thgencore/microservice-auth/internal/service/access"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/repository/access/converter"
	"github.com/8thgencore/microservice-auth/internal/repository/access/dao"
	"github.com/8thgencore/microservice-common/pkg/db"
)

const (
	tableName = "policies"

	endpointColumn     = "endpoint"
	allowedRolesColumn = "allowed_roles"
)

type repo struct {
	db db.Client
}

// NewRepository creates new object of repository layer.
func NewRepository(db db.Client) repository.AccessRepository {
	return &repo{db: db}
}

func (r *repo) GetRoleEndpoints(ctx context.Context) ([]*model.EndpointPermissions, error) {
	builderSelect := sq.Select(endpointColumn, allowedRolesColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "access_repository.GetRoleEndpoints",
		QueryRaw: query,
	}

	var endpointPermissions []*dao.EndpointPermissions
	err = r.db.DB().ScanAllContext(ctx, &endpointPermissions, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToEndpointPermissionsFromRepo(endpointPermissions), nil
}

func (r *repo) AddRoleEndpoint(ctx context.Context, endpoint string, allowedRoles []string) error {
	builderInsert := sq.Insert(tableName).
		Columns(endpointColumn, allowedRolesColumn).
		Values(endpoint, allowedRoles).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "access_repository.AddRoleEndpoint",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return accessService.ErrEndpointAlreadyExists
		}
	}

	return err
}

func (r *repo) UpdateRoleEndpoint(ctx context.Context, endpoint string, allowedRoles []string) error {
	builderUpdate := sq.Update(tableName).
		Set(allowedRolesColumn, allowedRoles).
		Where(sq.Eq{endpointColumn: endpoint}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "access_repository.UpdateRoleEndpoint",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)

	return err
}

func (r *repo) DeleteRoleEndpoint(ctx context.Context, endpoint string) error {
	builderDelete := sq.Delete(tableName).
		Where(sq.Eq{endpointColumn: endpoint}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "access_repository.DeleteRoleEndpoint",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)

	return err
}
