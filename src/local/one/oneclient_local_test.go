package main

import (
	"testing"
	"os"
	"log"
	coap "github.com/lorrin/go-coap"
)

func BenchmarkLocalOne(b *testing.B){
	req := coap.Message{
			Type:		coap.Confirmable,
			Code:		coap.GET,
			MessageID:	uint16(1),
			Payload:	[]byte("hello"),
	}
	path := "my/test"
	if len(os.Args) >1 {
			path = os.Args[1]
	}
	req.SetPathString(path)
	
	c, err := coap.Dial("udp","localhost:5683")
	if err != nil{
		log.Println("err in Dial..")
	}
	for i:=0;i<b.N;i++{
		rv, err := c.Send(req)
		if err != nil{
			if rv != nil{
				if err!= nil{
					log.Println("err in send..",err)
				}
//				payload := string(rv.Payload)
//				log.Println("Got response message payload:",payload,i)
			}
			rv, err = c.Receive()
		}
		
	}
}