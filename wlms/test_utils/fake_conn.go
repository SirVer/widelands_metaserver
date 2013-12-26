package test_utils

import (
	"io"
	. "launchpad.net/gocheck"
	"log"
	"net"
	"time"
)

type FakeConn struct {
	sendData_Reader *io.PipeReader
	sendData_Writer *io.PipeWriter
	recvData_Reader *io.PipeReader
	recvData_Writer *io.PipeWriter

	gotClosed *bool
}

func NewFakeConn(c *C) FakeConn {
	f := FakeConn{gotClosed: new(bool)}
	f.sendData_Reader, f.sendData_Writer = io.Pipe()
	f.recvData_Reader, f.recvData_Writer = io.Pipe()
	return f
}

func (f FakeConn) ServerWriter() io.Writer {
	return f.sendData_Writer
}

func (f FakeConn) ServerReader() io.Reader {
	return f.recvData_Reader
}

func (f FakeConn) GotClosed() bool {
	return *f.gotClosed
}

func (f FakeConn) Read(b []byte) (int, error) {
	n, err := f.sendData_Reader.Read(b)
	return n, err
}

func (f FakeConn) Write(b []byte) (n int, err error) {
	return f.recvData_Writer.Write(b)
}

func (f FakeConn) Close() error {
	f.sendData_Reader.Close()
	f.sendData_Writer.Close()
	f.recvData_Reader.Close()
	f.recvData_Writer.Close()
	*f.gotClosed = true
	return nil
}
func (f FakeConn) LocalAddr() net.Addr {
	return FakeAddr{}
}
func (f FakeConn) RemoteAddr() net.Addr {
	return FakeAddr{}
}
func (f FakeConn) SetDeadline(t time.Time) error {
	log.Print("Setting deadline %v", t)
	// NOCOM(sirver): implement
	return nil
}
func (f FakeConn) SetReadDeadline(t time.Time) error {
	// NOCOM(sirver): implement
	return nil
}
func (f FakeConn) SetWriteDeadline(t time.Time) error {
	// NOCOM(sirver): implement
	return nil
}