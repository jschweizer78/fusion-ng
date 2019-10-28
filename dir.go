package main

import (
	"os"
	"runtime"
)

var appDataRoot = "jrs"
var appDataDir = "fusion-editor"

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	appDataRoot = "." + appDataRoot
	return os.Getenv("HOME")
}
