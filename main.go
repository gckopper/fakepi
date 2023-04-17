package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Path map[string]struct {
	Method `yaml:",inline"`
}

type Method map[string]struct {
	File    string
	Body    string            `yaml:",omitempty"`
	Headers map[string]string `yaml:",omitempty"`
}

var data Path

func main() {
	logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	log.SetOutput(logFile)
	defer func(logFile *os.File) {
		err = logFile.Close()
		if err != nil {
			return
		}
	}(logFile)

	ip := flag.String("ip", "localhost", "The ip addrs this server will try to bind to")
	configFile := flag.String("config", "config.yml", "Path to config file")
	port := flag.Int("port", 8000, "Port this server will try to bind to")
	flag.Parse()

	yfile, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	data = make(Path)
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	http.Handle("/", http.HandlerFunc(handler))
	// Listen on localhost as this service should not be public
	addr := fmt.Sprint(*ip, ":", *port)
	fmt.Println("http://" + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		err = log.Output(0, fmt.Sprintln(err))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	method, exists := data[r.URL.RequestURI()]
	if !exists {
		log.Printf("This path is not allowed %s\n", r.URL.RequestURI())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	page, exists := method.Method[r.Method]
	if !exists {
		log.Printf("This method is not allowed %s\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	for k, v := range page.Headers {
		header := r.Header.Get(k)
		if header != v {
			log.Printf("Missing header %s: %s\n", k, v)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body\n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := string(b)
	if page.Body != body {
		log.Printf("This body is not allowed %s\n", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, page.File)
}
