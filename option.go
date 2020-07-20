package cuto

import "time"

type Option func(b *Browser)

func Binary(bin string) Option {
	return func(b *Browser) {
		b.binary = bin
	}
}

func Headless() Option {
	return func(b *Browser) {
		b.commands = append(b.commands, "--headless")
	}
}

func DataDir(dir string) Option {
	return func(b *Browser) {
		b.dataDir = dir
	}
}

func RemoteAddr(addr string) Option {
	return func(b *Browser) {
		b.remoteAddr = addr
	}
}

func Timeout(d time.Duration) Option {
	return func(b *Browser) {
		b.timeout = d
	}
}

func Debug() Option {
	return func(b *Browser) {
		b.debug = true
	}
}
