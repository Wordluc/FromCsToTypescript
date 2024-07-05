package main

import (
	"GoFromCsToTypescript/Writer"
	"os"
)

func main() {
	In:=os.Stdin
	Out:=os.Stdout
	Err:=os.Stderr
	str:=make([]byte,1024)
	In.Read(str)
	converted, err := Writer.Convert(string(str))
	if err != nil {
		Err.Write([]byte(err.Error()))
	}
	Out.Write([]byte(converted))
}
