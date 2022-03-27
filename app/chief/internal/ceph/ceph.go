package ceph

import (
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"github.com/google/wire"
)

var ProviderCephConnSet = wire.NewSet(NewCephConnect)

func NewCephConnect() *rados.Conn {
	conn, err := rados.NewConn()
	if err != nil {
		fmt.Println("error when invoke a new connection:", err)
	}
	err = conn.ReadDefaultConfigFile()
	err = conn.Connect()
	fmt.Println("connect ceph cluster ok!")
	return conn
}
