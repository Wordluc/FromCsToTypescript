package main

import (
	"GoFromCsToTypescript/Writer"
	"encoding/json"
	"errors"
	"net"
	"os"
)

type Response struct {
	Body   string
	Status Status
}
type Status struct {
	Code int
	Msg  string
}

func sendErrorResponse(e error, server net.PacketConn, addr net.Addr) {
	status := Status{Code: 500, Msg: e.Error()}
	resp := Response{Status: status}
	strResp, _ := json.Marshal(resp)
	server.WriteTo(strResp, addr)
}
func main() {
	port := os.Args[1]
	server, _ := net.ListenPacket("udp", "127.0.0.1:"+port)
	defer server.Close()
	str := make([]byte, 1000)
	resp := Response{}
	n, addr, e := server.ReadFrom(str)
	if e != nil {
		sendErrorResponse(e, server, addr)
		return ;
	}
	if n == 0 {
		sendErrorResponse(errors.New("no code to convert"), server, addr)
		return ;
	}
	strConverted, e := Writer.Convert(string(str))
	if e != nil {
		sendErrorResponse(e, server, addr)
		return ;
	}
	resp.Body = string(strConverted)
	resp.Status.Code = 200
	strResp, _ := json.Marshal(resp)
	server.WriteTo(strResp, addr)
}
