package main

import (
	"log"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
)

type TimeArgs struct{}

type WordReply struct {
	Word string
}

type TimeService struct{}

func (h *TimeService) Now(r *http.Request, args *TimeArgs, reply *WordReply) error {
	reply.Word = "HOGE"
	return nil
}

func main() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(TimeService), "")
	http.Handle("/", RPC)

	log.Println("Starting XML-RPC server on localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
