// +build darwin

package chr

import "os"

var (
	// Mac下默认浏览器
	defaultBrowser = []string{
		`/Applications/360Chrome.app/Contents/MacOS/360Chrome`,
		`/Applications/Google Chrome.app/Contents/MacOS/Google`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/chr"
)
