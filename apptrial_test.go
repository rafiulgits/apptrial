package apptrial_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/rafiulgits/apptrial"
)

func Test_HTTPServer(t *testing.T) {

	trail := apptrial.NewAppTrial("3STunnel", time.Minute*1, "_this_is_my_encrypt_key_")
	trail.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server is listening 8090")
	http.ListenAndServe(":8090", nil)
}
