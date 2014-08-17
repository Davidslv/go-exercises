package main

import (
  "bufio"
  "net"
  . "fmt"
)

var CRLF = ([]byte)("\n")
var i int = 0

func main() {
  for i < 3000000 {
    if address, e := net.ResolveUDPAddr("udp", ":1024"); e == nil {

      if server, e := net.DialUDP("udp", nil, address); e == nil {
        defer server.Close()

        if _, e = server.Write(CRLF); e == nil {
          if text, e := bufio.NewReader(server).ReadString('\n'); e == nil {
            Printf("%v", text)
          }
        }
      }
      i++
    }
  }

}
