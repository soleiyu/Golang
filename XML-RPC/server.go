package main

import (
	"log"
	"net/http"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
)

type Reply struct {
	Rep int
}

type Args struct {
	Vals []int
}

type Funcs struct {}

func (h *Funcs) GetSum(r *http.Request, args *Args, reply *Reply) error {

	sum := 0

	for _, val := range args.Vals {
		sum += val
	}

	reply.Rep = sum

	return nil
}

func main() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(Funcs), "")
	http.Handle("/", RPC)

	log.Println("Starting XML-RPC server on localhost:8000")
	http.ListenAndServe(":8000", nil)
}

