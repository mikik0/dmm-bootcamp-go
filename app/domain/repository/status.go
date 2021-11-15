package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindByID(ctx context.Context, id int64) (*object.Status, error)
	CreateStatus(ctx context.Context, status *object.Status) error
}
