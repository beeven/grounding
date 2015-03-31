package main

import(
	"encoding/xml"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"bytes"
)

type UploadFileInfo struct {
	EntNo		string	`xml:"entNo"`
	Filename	string	`xml:"filename"`
	Content		string 	`xml:"content"`
}


func DecodeXml(xmlDoc []byte, info* UploadFileInfo) {
	decoder := xml.NewDecoder(bytes.NewReader(xmlDoc))
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if startElement.Name.Local == "UploadFile" {
				decoder.DecodeElement(info, &startElement)
				return
			}
		}
	}
}


func main() {
	dat,err := ioutil.ReadFile("postData.xml")
	if err != nil {
		panic(err)
	}

	var uploadFile UploadFileInfo
	DecodeXml(dat,&uploadFile)
	body,err := base64.StdEncoding.DecodeString(uploadFile.Content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n",body)
}