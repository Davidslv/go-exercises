// Code example from "A Go Developer's Notebook"
// by @feyeleanor
// https://leanpub.com/GoNotebook

// To generate a certificate and a public key
// go run $GOROOT/src/pkg/crypto/tls/generate_cert.go -ca=true -host="localhost"


package main

import (
  . "fmt"
  . "net/http"
  "os"
  "sync"
)

var (
  address string
  secure_address string
  certificate string
  key string
)

var servers sync.WaitGroup

func init() {
  // $ SERVE_HTTP=":3030" go run web-hello-world
  if address = os.Getenv("SERVE_HTTP"); address == "" {
    address = ":1024"
  }

  if secure_address = os.Getenv("SERVE_HTTPS"); secure_address == "" {
    secure_address = ":1025"
  }

  if certificate = os.Getenv("SERVE_CERT"); certificate == "" {
    certificate = "cert.pem"
  }

  if key = os.Getenv("SERVE_KEY"); key == "" {
    key = "key.pem"
  }
}

func main() {
  message := "Hello David"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })

  Launch(func() {
    ListenAndServe(address, nil)
    Println("fuck SSL")
  })

  Launch(func() {
    ListenAndServeTLS(secure_address, certificate, key, nil)
    Println("under SSL")
  })

  servers.Wait()
}

func Launch(f func()) {
  servers.Add(1)

  go func() {
    defer servers.Done()
    f()
  }()
}
