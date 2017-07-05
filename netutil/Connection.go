package netutil

import (
	"io"
	"net"
	"time"
)

type Connection interface {
	io.ReadWriteCloser
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
	SetWriteDeadline(time.Time) error
	SetReadDeadline(time.Time) error
}
