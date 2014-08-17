package main

import(
  . "fmt"
  "net"
)

var HELLO_WORLD = ([]byte)("Hello David a tua mae suka!\n")

func main() {
  if address, e := net.ResolveUDPAddr("udp", ":1024"); e == nil {
    if server, e := net.ListenUDP("udp", address); e == nil {
      i := 0
      for buffer := MakeBuffer(); ; buffer = MakeBuffer() {
        if n, client, e := server.ReadFromUDP(buffer); e == nil {

          go func(c *net.UDPAddr, packet []byte) {
            if n, e := server.WriteToUDP(HELLO_WORLD, c); e == nil {
              Printf("%v bytes written to: %v - %v\n", n, c, i)
              i++
            }
          }(client, buffer[:n])
        }
      }
    }
  }
}

func MakeBuffer() (r []byte) {
  return make([]byte, 1024)
}
