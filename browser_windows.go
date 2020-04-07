// +build windows

package cuto

import "os"

var (
	// Windows下默认浏览器
	defaultBrowser = []string{
		`C:\Program Files (x86)\Google\cutoome\Application\cutoome.exe`,
		`D:\Program Files (x86)\Google\cutoome\Application\cutoome.exe`,
		`C:\Program Files\Google\cutoome\Application\cutoome.exe`,
		`D:\Program Files\Google\cutoome\Application\cutoome.exe`,
		os.Getenv("USERPROFILE") + `\AppData\Local\Google\cutoome\Application\cutoome.exe`,
		os.Getenv("USERPROFILE") + `\AppData\Roaming\360se6\Application\360se.exe`,
	}
	defaultUserDataTmpDir = os.TempDir() + `\cuto`
)
