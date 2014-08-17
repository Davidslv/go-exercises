package main

import (
  "bufio"
  "net"
  . "fmt"
)

var CRLF = ([]byte)("\n")


func main() {
  if address, e := net.ResolveUDPAddr("udp", ":1024"); e == nil {

    if server, e := net.DialUDP("udp", nil, address); e == nil {
      defer server.Close()
      for b := 0; b < 3; b++ {
        if _, e = server.Write(CRLF); e == nil {
          if text, e := bufio.NewReader(server).ReadString('\n'); e == nil {
            Printf("%v number: %v, connection: %v", text, i, b)
          }
        }
      }
    }
  }
}
