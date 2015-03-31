package main

import (
    "bytes"
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    url := "http://localhost:9001/grounding/12345"
    fmt.Println("URL:>", url)

    var postData = []byte(`<?xml version="1.0" encoding="utf-8"?><root><element></element></root>`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
    req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
    req.Header.Set("Content-Length", len(postData))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
    