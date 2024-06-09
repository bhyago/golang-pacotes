package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "John"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

	//c := http.Client{Timeout: time.Duration(1) * time.Second}
	//resp, err := c.Get("http://google.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//println(string(body))
}
