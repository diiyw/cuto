// +build windows

package cuto

import "os"

var (
	// Windows下默认浏览器
	defaultBrowser = []string{
		`C:\Program Files (x86)\Google\chrome\Application\chrome.exe`,
		`D:\Program Files (x86)\Google\chrome\Application\chrome.exe`,
		`C:\Program Files\Google\chrome\Application\chrome.exe`,
		`D:\Program Files\Google\chrome\Application\chrome.exe`,
		os.Getenv("USERPROFILE") + `\AppData\Local\Google\chrome\Application\chrome.exe`,
		os.Getenv("USERPROFILE") + `\AppData\Roaming\360se6\Application\360se.exe`,
	}
	defaultUserDataTmpDir = os.TempDir() + `\cuto`
)
