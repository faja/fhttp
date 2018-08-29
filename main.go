package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

var f = flag.String("f", "httpContent.yaml", "path to static file to serve")
var p = flag.Int("p", 8042, "port http server listens on")

type response struct {
	Code    int    `yaml:"code"`
	Content string `yaml:"content"`
}

func main() {
	flag.Parse()

	resp := response{
		Code: 200,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadFile(*f)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}

		err = yaml.Unmarshal(d, &resp)
		if err != nil {
			log.Fatalf("could not unmarshal yaml: %v", err)
		}

		w.WriteHeader(resp.Code)
		fmt.Fprintf(w, resp.Content)
	})
	addr := fmt.Sprintf("127.0.0.1:%d", *p)
	log.Printf("Starting http server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
