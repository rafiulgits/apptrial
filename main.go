package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	time_format = "2006-01-02T15:04:05.000Z"
)

type AppTrial struct {
	AppName       string
	EncryptionKey string
	Duration      time.Duration
}

/*
NewAppTrial : key size must be 16/24/32bytes
*/
func NewAppTrial(appName string, duration time.Duration, key string) *AppTrial {
	return &AppTrial{
		AppName:       appName,
		Duration:      duration,
		EncryptionKey: key,
	}
}

func (appTrial *AppTrial) Start() {
	go appTrial.checker()
}

func (appTrial *AppTrial) checker() {
	if !isFileExist(appTrial.AppName) {
		appTrial.saveState()
	}

	filePath := getFileLocation(appTrial.AppName)
	fileData, err := ReadFromFile(filePath)
	if err != nil {
		panic(err)
	}
	data := Decrypt(string(fileData), appTrial.EncryptionKey)
	expiredTime, _ := time.Parse(time_format, data)
	for {
		if appTrial.isExpired(expiredTime) {
			panic("Application is expired")
		}
		time.Sleep(time.Second * 2)
	}
}

func (appTrial *AppTrial) saveState() {
	dir := getRootDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.Mkdir(dir, os.ModePerm)
	}
	expiredTime := time.Now().Add(appTrial.Duration).UTC().Format(time_format)
	data := Encrypt(expiredTime, appTrial.EncryptionKey)
	filePath := getFileLocation(appTrial.AppName)
	WriteToFile(data, filePath)
}

func (appTrial *AppTrial) isExpired(expiredTime time.Time) bool {
	return time.Now().UTC().After(expiredTime)
}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server is listening 8090")
	http.ListenAndServe(":8090", nil)
}

func main() {
	trail := NewAppTrial("3STunnel", time.Minute*2, "_this_is_my_encrypt_key_")
	trail.Start()
	httpServer()
}
