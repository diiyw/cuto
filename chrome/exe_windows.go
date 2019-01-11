// +build windows

package chrome

import "os"

var (
	// Windows下默认浏览器
	defaultChrome = []string{
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
		`D:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files\Google\Chrome\Application\chrome.exe`,
		`D:\Program Files\Google\Chrome\Application\chrome.exe`,
		os.Getenv("USERPROFILE") + `\AppData\Local\Google\Chrome\Application\chrome.exe`,
		os.Getenv("USERPROFILE") + `\AppData\Roaming\360se6\Application\360se.exe`,
	}
	defaultUserDataTmpDir = os.TempDir()
)
