package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindByID : IDで検索
func (r *status) FindByID(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// CreateStatus : ステータスを作成
func (r *status) CreateStatus(ctx context.Context, status *object.Status) error {
	_, err := r.db.NamedExecContext(ctx, "insert into status (account_id, content, create_at) values (:account_id, :content, :created_at, :updated_at)", status)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
