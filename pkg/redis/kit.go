package redis

import "github.com/go-redis/redis/v8"

const (
	// 比较删除lua脚本
	compareAndDeleteScript = `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
    	return redis.call("DEL", KEYS[1])
	else
    	return 0
	end
	`
)

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
