package apptrial

import (
	"os"
	"path"
	"runtime"
)

const (
	file_name = "exp.d.encrpt"
)

func getRootDir() string {
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

func getPath(folder string) string {
	rootPath := getRootDir()
	path := path.Join(rootPath, folder)
	return path
}

func getFileLocation(appName string) string {
	folderPath := getPath(appName)
	return path.Join(folderPath, file_name)
}

func isFileExist(appName string) bool {
	filePath := getFileLocation(appName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
