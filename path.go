package main

import (
	"os"
	"path"
	"runtime"
)

func getLocalPath() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return path.Join(home, "AppData")
	} else if runtime.GOOS == "darwin" {
		//Mac OSX
		return "/tmp"
	}
	return os.TempDir()
}

func GetPath(folder string) string {
	rootPath := getLocalPath()
	path := path.Join(rootPath, folder)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
	return path
}
