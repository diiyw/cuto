// +build darwin

package chrome

import "os"

var (
	// Mac下默认浏览器
	defaultChrome = []string{
		`/Applications/360Chrome.app/Contents/MacOS/360Chrome`,
		`/Applications/Google Chrome.app/Contents/MacOS/Google`,
		`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/goc"
)
