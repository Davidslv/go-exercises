package main

import (
  "crypto/rand"
  "crypto/tls"
  . "fmt"
)

func main() {
  if certificate, e := tls.LoadX509KeyPair("cert.server.pem", "key.server.pem"); e == nil {
    config := tls.Config {
      Certificates: []tls.Certificate { certificate },
      Rand: rand.Reader,
    }

    if listener, e := tls.Listen("tcp", ":1025", &config); e == nil {
      i := 0
      for {
        if connection, e := listener.Accept(); e == nil {
          go func(c *tls.Conn) {
            defer c.Close()
            i++
            Fprintln(c, "hello Worldx", i)
          }(connection.(*tls.Conn))
        }
      }
    }
  }
}
