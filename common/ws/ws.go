package ws

import (
	"fmt"
	"github.com/mars-projects/mars/api"
)

func NewWs() *Ws {
	return &Ws{}
}

type Ws struct {
}

func (e *Ws) Send(message *api.Reply) error {
	fmt.Println(message)
	return nil
}
