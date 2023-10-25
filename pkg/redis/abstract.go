package redis

import "context"

type Kiter interface {
	Lock(ctx context.Context) (bool, error)
	UnLock(ctx context.Context) error
}
