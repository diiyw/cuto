// +build darwin

package cuto

import "os"

var (
	// Mac下默认浏览器
	defaultBrowser = []string{
		`/Applications/360Chrome.app/Contents/MacOS/360Chrome`,
		`/Applications/Google Chrome.app/Contents/MacOS/Google`,
		`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`,
		`/Applications/360cutoome.app/Contents/MacOS/360cutoome`,
		`/Applications/Google cutoome.app/Contents/MacOS/Google`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/cuto"
)
