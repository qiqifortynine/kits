package alter

type Alerter interface {
	Send() error
}
