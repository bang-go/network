package udpx

import (
	"github.com/bang-go/opt"
	"time"
)

type connectOptions struct {
	timeout time.Duration
}

func WithConnectTimeout(timeout time.Duration) opt.Option[connectOptions] {
	return opt.OptionFunc[connectOptions](func(o *connectOptions) {
		o.timeout = timeout
	})
}
