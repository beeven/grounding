package main

import(
	"fmt"
	"encoding/base64"
	"io/ioutil"
)

var template string = `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http:/www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
<soap12:Body>
    <UploadFile xmlns="http://tempuri.org/">
        <entNo>PTE51001407270000001</entNo>
        <filename>test.xml</filename>
        <content>%v</content>
    </UploadFile>
</soap12:Body>
</soap12:Envelope>`

func main() {
	dat,err := ioutil.ReadFile("message.xml")
	if err != nil {
		panic(err);
	}
	encoded := base64.StdEncoding.EncodeToString(dat)
	output := fmt.Sprintf(template, encoded)
	ioutil.WriteFile("postData.xml",[]byte(output), 0644)
}