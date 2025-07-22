package ziface

import "net"

type IConnection interface {
	Start()

	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32
}

type HandFunc func(*net.TCPConn, []byte, int) error
