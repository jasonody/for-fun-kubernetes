// http -f POST :8000/save hello=jason
// http ':8000/read?key=hello'

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		os.Mkdir("./data", 0700)
	}

	http.HandleFunc("/save", save)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8000", nil)
}

func save(rw http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body)
	keyValuePair := strings.Split(string(data), "=")
	key, value := keyValuePair[0], keyValuePair[1]

	filename := composeFilename(key)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to open file %q due to: %s", filename, err.Error())
		outputMessage(rw, errMsg)
		return
	}
	defer file.Close()

	if _, err = file.WriteString(value); err != nil {
		errMsg := fmt.Sprintf("Failed to save data %q due to: %s", string(data), err.Error())
		outputMessage(rw, errMsg)
		return
	}

	successMsg := fmt.Sprintf("Saved value %q to %q", value, filename)
	outputMessage(rw, successMsg)
}

func read(rw http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	filename := composeFilename(key)
	value, _ := ioutil.ReadFile(filename)
	if string(value) == "" {
		errMsg := fmt.Sprintf("Key %q does not exist in the store", key)
		outputMessage(rw, errMsg)
		return
	}

	successMsg := fmt.Sprintf("Value for key %q is %q", key, string(value))
	outputMessage(rw, successMsg)
}

func composeFilename(key string) string {
	return fmt.Sprintf("./data/%s", key)
}

func outputMessage(rw http.ResponseWriter, msg string) {
	rw.Write([]byte(msg))
	fmt.Println(msg)
}
