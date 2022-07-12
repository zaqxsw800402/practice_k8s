package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" {
		fmt.Println("HOST not given. exiting")
		return
	}
	for {
		postTo6666(host, port)
		fmt.Println(getLastLineWithSeek("tmp/log.txt"))
		time.Sleep(time.Second * 1)
	}
}

func postTo6666(host, port string) {
	req, err := http.NewRequest(http.MethodGet, "http://"+host+":"+port, http.NoBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func getLastLineWithSeek(filepath string) string {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		panic("Cannot open file")
		os.Exit(1)
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}
