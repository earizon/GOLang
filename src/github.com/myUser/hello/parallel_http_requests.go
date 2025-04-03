package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	URL_TEST = "https://httpbin.org/get"
)

// [[{io.http]]
func parallel_requests() {
	/**
	 * No need for a worker pools. The net.http client takes care of
	 * it all. 
	 * WARN: This also mean extra care to avoid "too many" parallel 
	 *    client connections. A server to server code will probably
	 *    need some barrier to lock down how many parallel requests 
	 *    can be made to this function.
	 *
	 */
	resp, err := http.Get(URL_TEST)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println("body:" + string(body[:]))
}
// [[io.http}]]


func main() {
	for i := 0; i < 1000; i++ {
		go A()
	}
	time.Sleep(10 * time.Second)
}
