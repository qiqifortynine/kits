package redis

type Kiter interface {
	Lock() error
	UnLock() error
}
