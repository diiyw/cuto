// +build darwin

package cuto

import "os"

var (
	// Mac下默认浏览器
	defaultBrowser = []string{
		`/Applications/360cutoome.app/Contents/MacOS/360cutoome`,
		`/Applications/Google cutoome.app/Contents/MacOS/Google`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/cuto"
)
