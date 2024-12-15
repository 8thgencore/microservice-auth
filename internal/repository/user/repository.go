package user

import (
	"context"
	"errors"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/repository/user/converter"
	"github.com/8thgencore/microservice-auth/internal/repository/user/dao"
	userService "github.com/8thgencore/microservice-auth/internal/service/user"

	"github.com/8thgencore/microservice-common/pkg/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	passwordColumn  = "password"
	emailColumn     = "email"
	roleColumn      = "role"
	versionColumn   = "version"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	userNameKey  = "users_name_key"
	userEmailKey = "users_email_key"
)

type repo struct {
	db db.Client
}

// NewRepository creates new object of repository layer.
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.UserCreate) (string, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, nameColumn, roleColumn, emailColumn, passwordColumn).
		Values(user.ID, user.Name, user.Role, user.Email, user.Password).
		Suffix("RETURNING " + idColumn)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			switch pgErr.ConstraintName {
			case userNameKey:
				return "", userService.ErrUserNameExists
			case userEmailKey:
				return "", userService.ErrUserEmailExists
			}
		}

		return "", err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id string) (*model.User, error) {
	builderSelect := sq.Select(
		idColumn,
		nameColumn,
		emailColumn,
		roleColumn,
		versionColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user dao.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, userService.ErrUserNotFound
		}

		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, user *model.UserUpdate) error {
	builderUpdate := sq.Update(tableName).
		Set(updatedAtColumn, sq.Expr("NOW()")).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	if user.Name.Valid {
		builderUpdate = builderUpdate.Set(nameColumn, user.Name.String)
	}
	if user.Email.Valid {
		builderUpdate = builderUpdate.Set(emailColumn, user.Email.String)
	}
	if user.Role.Valid {
		builderUpdate = builderUpdate.Set(roleColumn, user.Role.String)
	}
	if user.Version.Valid {
		builderUpdate = builderUpdate.Set(versionColumn, user.Version.Int32)
	}

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			switch pgErr.ConstraintName {
			case userNameKey:
				return userService.ErrUserNameExists
			case userEmailKey:
				return userService.ErrUserEmailExists
			}
		}

		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAuthInfo(ctx context.Context, username string) (*model.AuthInfo, error) {
	builderSelect := sq.Select(idColumn, nameColumn, roleColumn, passwordColumn, versionColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{nameColumn: username}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetAuthInfo",
		QueryRaw: query,
	}

	var authInfo dao.AuthInfo
	err = r.db.DB().ScanOneContext(ctx, &authInfo, q, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return converter.ToAuthInfoFromRepo(&authInfo), nil
}

// FindByName returns user with specified name
func (r *repo) FindByName(ctx context.Context, name string) (*model.User, error) {
	builderSelect := sq.Select(
		idColumn,
		nameColumn,
		emailColumn,
		roleColumn,
		versionColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{nameColumn: name}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.FindByName",
		QueryRaw: query,
	}

	var user dao.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
