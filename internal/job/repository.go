package job

import "context"

type Repository interface {
	Create (ctx context.Context, job Job) error;
	GetById (ctx context.Context, id string) (*Job, error);
	ListActive (ctx context.Context) ([]Job, error);
	UpdateState (ctx context.Context, id string, state string) error;
}