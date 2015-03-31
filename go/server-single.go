package main

import (
	"fmt"
	"net/http"
	"runtime"
	"log"
	"strings"
	"go-uuid/uuid"
	"os"
	"io"
)

type GroundingServer struct {}

func PipeRequest(r *http.Request, path string) {
	if(r.Body != nil) {
		f,err := os.Create(path)
		if err == nil {
			io.Copy(f, r.Body)
			f.Close()
		} else {
			log.Fatal(err)
		}
	}
}

func (gs GroundingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL)

	path := r.URL.Path
	var compolents []string
	//var params map[string]string
	//params := r.URL.Query()
	strs:=strings.Split(path,"/")

	for i := range strs {
		s := strs[i]
		if s != "" {
			compolents = append(compolents,s)
		}
	}
	//fmt.Println(compolents)
	//fmt.Println(params)

	fileid := uuid.NewRandom().String()

	//fmt.Println(fileid)

	if(compolents[0] != "grounding" || len(compolents[1]) != 5 || len(compolents) > 2) {
		
		w.WriteHeader(402);
		w.Write([]byte("Bad request"));
		return
	}

	filepath := "../data/" + fileid + ".xml"
	PipeRequest(r, filepath);
	w.Write([]byte("Message received."));
}

func main() {
	runtime.GOMAXPROCS(1)

	var gs GroundingServer
    log.Fatal(http.ListenAndServe("172.7.1.16:9001",gs))
    fmt.Println("Listening on port: 9001");
}
