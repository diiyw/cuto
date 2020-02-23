package chr

import "time"

type Option func(b *Browser)

func WithChrome(bin string) Option {
	return func(b *Browser) {
		b.bin = bin
	}
}

func WithDataDir(dir string) Option {
	return func(b *Browser) {
		b.dataDir = dir
	}
}

func WithRemoteAddr(addr string) Option {
	return func(b *Browser) {
		b.remoteAddr = addr
	}
}

func WithTimeout(d time.Duration) Option {
	return func(b *Browser) {
		b.timeout = d
	}
}
