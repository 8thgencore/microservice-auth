package access

import (
	"context"

	sq "github.com/Masterminds/squirrel"

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
