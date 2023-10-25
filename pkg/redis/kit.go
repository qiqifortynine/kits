package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

/**
	redis锁工具
**/

const (
	// 比较删除lua脚本
	compareAndDeleteScript = `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
    	return redis.call("DEL", KEYS[1])
	else
    	return 0
	end
	`
	defaultExp = time.Second * 10
)

func NewKiter(client *redis.Client, key string) Kiter {
	return &Kit{Client: client, Key: key, uuid: uuid.New().String()}
}

type Kit struct {
	Client     *redis.Client
	Key        string
	uuid       string
	cancelFunc context.CancelFunc
}

func (k *Kit) Lock(ctx context.Context) (suc bool, err error) {
	suc, err = k.Client.SetNX(ctx, k.Key, k.uuid, defaultExp).Result()
	if err != nil || !suc {
		return false, err
	}

	c, cancel := context.WithCancel(ctx)
	k.cancelFunc = cancel
	k.refresh(c)

	return suc, nil
}

func (k *Kit) UnLock(ctx context.Context) error {
	resp, err := k.Client.Eval(ctx, compareAndDeleteScript, []string{k.Key}, k.uuid).Result()
	if err != nil {
		return err
	}
	if resp != 0 {
		k.cancelFunc()
	}

	return nil
}

func (k *Kit) refresh(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(defaultExp / 4)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				k.Client.Expire(ctx, k.Key, defaultExp)
			}
		}
	}()
}
