package execution

import "context"

type Repository interface {
	Create (ctx context.Context, exec JobExecution) error;
	GetById (ctx context.Context, id string) (*JobExecution, error);
	UpdateState (ctx context.Context, id string, state string, errMsg *string) error;
	ListPendingRetries (ctx context.Context) ([]JobExecution, error);
}