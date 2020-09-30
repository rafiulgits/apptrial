package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

const (
	FILE_NAME   = "exp.conf"
	TIME_FORMAT = "2006-01-02T15:04:05.000Z"
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

func IsExpired(startedTime time.Time) bool {
	whenToExpired := startedTime.Add(time.Second * 300)
	return whenToExpired.After(time.Now())
}

func main() {
	key := "testtesttesttest"
	if !IsFileExist() {
		nowTime := time.Now().Format(TIME_FORMAT)
		fmt.Println(nowTime)
		data := Encrypt(nowTime, key)
		filePath := GetFilePath()
		fmt.Println(filePath)
		WriteToFile(data, filePath)
		fmt.Println("A New File Written")
	} else {
		filePath := GetFilePath()
		fileData, err := ReadFromFile(filePath)
		if err != nil {
			fmt.Println(err.Error())
		}
		data := Decrypt(string(fileData), key)
		startedTime, _ := time.Parse(TIME_FORMAT, data)
		fmt.Println("Time: ", startedTime.String())
		if IsExpired(startedTime) {
			panic("Application is expired")
		}
		fmt.Println(data)
	}
}
