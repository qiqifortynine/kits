package redis

import "github.com/go-redis/redis/v8"

func NewKiter(client *redis.Client) Kiter {
	return &Kit{client: client}
}

type Kit struct {
	client *redis.Client
}

func (k *Kit) Lock() (err error) {
	return nil
}

func (k *Kit) UnLock() (err error) {
	return nil
}
