package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mars-projects/mars/api"
	"log"
)

func main() {
	val := make([]byte, 2)
	val[0] = 0
	val[1] = 1
	genericRPC := api.GenericRPC{
		Method:     "register.app",
		From:       "gpio",
		To:         "coreservice",
		Code:       0,
		Parameters: val,
	}
	data, err := proto.Marshal(&genericRPC)
	if err != nil {
		log.Fatal("genericRPC marshaling error: ", err)
	}
	websocketMessage := api.WebsocketMessage{
		Topic: "weatherstation.RPC",
		Body:  data,
	}
	mes, err := proto.Marshal(&websocketMessage)
	if err != nil {
		log.Fatal("wesocketMessage marshaling error: ", err)
	}
	deWebsocketMessage := &api.WebsocketMessage{}
	err = proto.Unmarshal(mes, deWebsocketMessage)

	if err != nil {
		log.Fatal("websocketMessage unmarshaling error: ", err)
	}
	fmt.Printf("Topic=%s\n", deWebsocketMessage.Topic)
	deGenericPRC := &api.GenericRPC{}
	err = proto.Unmarshal(deWebsocketMessage.Body, deGenericPRC)

	if err != nil {
		log.Fatal("websocketMessage unmarshaling error: ", err)
	}
	fmt.Printf("Method=%s\n", deGenericPRC.Method)
}
