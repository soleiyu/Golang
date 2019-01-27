package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
)

type Args struct {
	Vals []int
}

type Reply struct {
	Rep int
}

func XmlRpcCall(method string, args Args) (reply Reply, err error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	body := bytes.NewBuffer(buf)

	resp, err := http.Post("http://192.168.2.106:8000/", "text/xml", body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	xml.DecodeClientResponse(resp.Body, &reply)
	return
}

func summation(num int) {

	args := Args{}
	args.Vals = make([]int, num)

	for i := 0; i < num; i++ {
		args.Vals[i] = i + 1
	}

	reply, _ := XmlRpcCall("Funcs.GetSu", args)
	fmt.Println(reply.Rep)
}

func main() {
	summation(4)
}
