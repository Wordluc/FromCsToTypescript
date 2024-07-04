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
func mappingErrorResponse(e error)Response{
	   var status=Status{Code:500,Msg:e.Error()}
		 return Response{Status:status}
}
func main() {
	port := os.Args[1]
	server, _ := net.ListenPacket("udp", "127.0.0.1:"+port)
	defer server.Close()
	str := make([]byte, 1000)
	resp:=Response{}
	n, addr, e := server.ReadFrom(str)
	if e != nil {
    resp=mappingErrorResponse(e)
	}
	if n == 0 {
    resp=mappingErrorResponse(errors.New("none code to convert"))
	}
	strConverted, e := Writer.Convert(string(str))
	if e != nil {
    resp=mappingErrorResponse(e)
	}
	resp.Body=string(strConverted)
	resp.Status.Code=200
  strResp,_:=json.Marshal(resp)
	server.WriteTo(strResp,addr)
}
