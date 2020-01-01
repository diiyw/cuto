// +build linux

package chrome

import "os"

var (
	// Linux下默认浏览器
	defaultChrome = []string{
		`/usr/bin/google-chrome`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/chr"
)
