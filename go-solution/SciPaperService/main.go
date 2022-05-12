package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var reply string

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTIzOTIwNjksImlhdCI6MTY1MjM4ODQ2OSwibmFtZSI6Im1ha2FuYWxiZXJ0MiJ9.Cl905LCPqgBk7YmMRL_2vfHwcIdIkkAbTaBQCfvYSHs"
	client.Call("RPC.GetName", token, &reply)

	fmt.Println(reply)
}