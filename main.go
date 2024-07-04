package main

import (
	"GoFromCsToTypescript/Writer"
	"encoding/json"
	"errors"
	"fmt"
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

func sendErrorResponse(e error, conn net.Conn) {
	status := Status{Code: 500, Msg: e.Error()}
	resp := Response{Status: status}
	strResp, _ := json.Marshal(resp)
	conn.Write(strResp)
}

func main() {
	port := os.Args[1]
	server, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer server.Close()
	fmt.Println("Server listening on port", port)
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}
		result:=make(chan bool)
		go handleConnection(conn,result)
		if <-result{
			break
		}

	}
}

func handleConnection(conn net.Conn,result chan bool) {
	defer conn.Close()
	str := make([]byte, 10000)
	n, err := conn.Read(str)
	if err != nil {
		sendErrorResponse(err, conn)
		result<-true
	}
	if string(str)=="close"{
		result<-true
	}
	if n == 0 {
		sendErrorResponse(errors.New("no code to convert"), conn)
		result<-true
	}
	strConverted, err := Writer.Convert(string(str[:n]))
	if err != nil {
		sendErrorResponse(err, conn)
		result<-true
	}
	resp := Response{
		Body: strConverted,
		Status: Status{
			Code: 200,
			Msg:  "Success",
		},
	}
	strResp, _ := json.Marshal(resp)
	conn.Write(strResp)
	result<-true
}

