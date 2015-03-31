package main

import (
	"fmt"
	"net/http"
	"runtime"
	"log"
	"strings"
	"encoding/xml"
	"encoding/base64"
	"bytes"
	"io/ioutil"
)

type GroundingServer struct {}


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




func PipeRequest(r *http.Request, path string) {
	if(r.Body != nil && r.Header.Get("Content-Type") == "application/soap+xml; charset=utf-8") {
		var uploadFile UploadFileInfo

		body,err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		DecodeXml(body,&uploadFile)
		content, err := base64.StdEncoding.DecodeString(uploadFile.Content)
		if err != nil {
			panic(err)
		}

		//fmt.Printf("%v : %q\n",uploadFile.Filename,content)

		err = ioutil.WriteFile(path + "/" + uploadFile.Filename, content, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func (gs GroundingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL)

	path := r.URL.Path
	var compolents []string

	strs:=strings.Split(path,"/")

	for i := range strs {
		s := strs[i]
		if s != "" {
			compolents = append(compolents,s)
		}
	}


	if(compolents[0] != "grounding" || len(compolents[1]) != 5 || len(compolents) > 2) {
		
		w.WriteHeader(402);
		w.Write([]byte("Bad request"));
		return
	}

	filepath := "../data" 
	PipeRequest(r, filepath);
	w.Write([]byte("Message received."));
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	var gs GroundingServer
    log.Fatal(http.ListenAndServe("localhost:9001",gs))
    fmt.Println("Listening on port: 9001");
}
