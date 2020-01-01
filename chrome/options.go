package chrome

import "time"

type Option func(chrome *Chrome)

func WithChrome(bin string) Option {
	return func(chrome *Chrome) {
		chrome.bin = bin
	}
}

func WithDataDir(dir string) Option {
	return func(chrome *Chrome) {
		chrome.dataDir = dir
	}
}

func WithRemoteAddr(addr string) Option {
	return func(chrome *Chrome) {
		chrome.remoteAddr = addr
	}
}

func WithTimeout(d time.Duration) Option {
	return func(chrome *Chrome) {
		chrome.timeout = d
	}
}
