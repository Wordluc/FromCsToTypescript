package main

import ("net")
func main() {
	n,_:=net.Dial("udp", "0.0.0.0:300")
	n.Write([]byte("hello"))
  defer n.Close()
}
