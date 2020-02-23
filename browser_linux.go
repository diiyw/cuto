// +build linux

package chr

import "os"

var (
	// Linux下默认浏览器
	defaultBrowser = []string{
		`/usr/bin/google-chrome`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/chr"
)
