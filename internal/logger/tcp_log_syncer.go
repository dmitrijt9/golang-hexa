package logger

import "net"

type TcpLogSyncer struct {
	conn net.Conn
}

func (wsSync *TcpLogSyncer) Write(p []byte) (n int, err error) {
	if wsSync.conn != nil {
		n, err = wsSync.conn.Write(p)
	}
	return
}

func (wsSync *TcpLogSyncer) Sync() error {
	// nothing to sync for tcp websocket
	return nil
}

func (wsSync *TcpLogSyncer) SetTcpConnection(conn net.Conn) {
	wsSync.conn = conn
}
