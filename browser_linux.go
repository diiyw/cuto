// +build linux

package cuto

import "os"

var (
	// Linux下默认浏览器
	defaultBrowser = []string{
		`/usr/bin/google-cutoome`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/cuto"
)
