package net

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync/atomic"
)

var cidBase uint64 = 10000

type WsConnection struct {
	Cid       string
	Conn      *websocket.Conn
	manager   *WsManager
	ReadChan  chan *MsgPack
	WriteChan chan []byte
	Session   *Session
}

func (c *WsConnection) Run() {

}

func (c *WsConnection) Close() {

}

func (c *WsConnection) SendMessage(buf []byte) error {
	return nil
}

func (c *WsConnection) GetSession() *Session {
	return nil
}

func NewWsConnection(conn *websocket.Conn, manager *WsManager) *WsConnection {
	cid := fmt.Sprintf("%s-%s-%d", uuid.New().String(), manager.ServerId, atomic.AddUint64(&cidBase, 1))
	return &WsConnection{
		Conn:      conn,
		manager:   manager,
		Cid:       cid,
		WriteChan: make(chan []byte, 1024),
		ReadChan:  manager.ClientReadChan,
		Session:   NewSession(cid),
	}
}
