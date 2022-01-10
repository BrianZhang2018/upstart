package mysql

import (
	"brian.com/upstart/pkg/model"
	"context"
)

type (
	SqlDb interface {
		Close() error
		Health(ctx context.Context) error
		GetPayments(request model.Customer)
	}
)
