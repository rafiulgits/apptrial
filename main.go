package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	FILE_NAME        = "exp.conf"
	TIME_FORMAT      = "2006-01-02T15:04:05.000Z"
	KEY              = "testtesttesttest"
	EXPIRED_DURATION = time.Minute * 2
)

func GetFilePath() string {
	rootPath := GetPath("3STunnel")
	return path.Join(rootPath, FILE_NAME)
}

func IsFileExist() bool {
	filePath := GetFilePath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsExpired(expiredTime time.Time) bool {
	return time.Now().UTC().After(expiredTime)

}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server is listening 8090")
	http.ListenAndServe(":8090", nil)
}

func saveState() {
	nowTime := time.Now().Add(EXPIRED_DURATION).UTC().Format(TIME_FORMAT)
	fmt.Println(nowTime)
	data := Encrypt(nowTime, KEY)
	filePath := GetFilePath()
	fmt.Println(filePath)
	WriteToFile(data, filePath)
	fmt.Println("A New File Written")
}

func expireChecker() {
	if !IsFileExist() {
		saveState()
	}
	filePath := GetFilePath()
	fileData, err := ReadFromFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	data := Decrypt(string(fileData), KEY)
	expiredTime, _ := time.Parse(TIME_FORMAT, data)
	for {
		if IsExpired(expiredTime) {
			panic("Application is expired")
		}
		time.Sleep(time.Second * 2)
	}

}

func main() {
	go expireChecker()
	httpServer()
}
