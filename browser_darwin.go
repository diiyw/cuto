// +build darwin

package cuto

import "os"

var (
	// Mac下默认浏览器
	defaultBrowser = []string{
		`/Applications/360chrome.app/Contents/MacOS/360chrome`,
		`/Applications/Google chrome.app/Contents/MacOS/Google`,
	}
	defaultUserDataTmpDir = os.TempDir() + "/cuto"
)
