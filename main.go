package main

import (
	"net"
	"os"
)
func main() {
	port:=os.Args[1]
	n,_:=net.Dial("udp", "0.0.0.0:"+port)
	n.Write([]byte("hello"))
  defer n.Close()
}
