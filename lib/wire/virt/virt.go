package virt

import (
	"github.com/google/wire"
	"libvirt.org/go/libvirt"
)

var ProviderLibvirtConnectSet = wire.NewSet(NewLibvirtConnect)

func NewLibvirtConnect() *libvirt.Connect {
	connect, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return nil
	}
	return connect
}
