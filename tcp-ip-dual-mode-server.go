package main

import(
  "crypto/rand"
  "crypto/tls"
  "net"
  "sync"
  . "fmt"
)

var servers sync.WaitGroup

func main() {
  if listener, e := net.Listen("tcp", ":1024"); e == nil {
    Serve(listener)
  }

  Serve(TLSListener("cert.server.pem", "key.server.pem", ":1025"))

  servers.Wait()
}

func TLSListener(cert, key, address string) (r net.Listener) {
  if certificate, e := tls.LoadX509KeyPair(cert, key); e == nil {
    config := tls.Config{
      Certificates: []tls.Certificate{ certificate },
      Rand: rand.Reader,
    }

    if listener, e := tls.Listen("tcp", address, &config); e == nil {
      r = listener
    }
  }
  return
}

func Serve(listener net.Listener) {
  if listener != nil {
    Launch(func() {
      i := 0
      for {
        if connection, e := listener.Accept(); e == nil {
          go func(c net.Conn) {
            defer c.Close()
            i++
            Fprintln(c, "hello stranger!", i)
          }(connection)
        }
      }
    })
  }
}

func Launch(f func()) {
  servers.Add(1)
  go func() {
    defer servers.Done()
    f()
  }()
}
