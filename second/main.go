package main

import (
	"fmt"
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
