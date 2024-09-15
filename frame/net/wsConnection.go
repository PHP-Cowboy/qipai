package net

import (
	"common/global"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

var (
	pongWait             = 10 * time.Second
	writeWait            = 10 * time.Second
	pingInterval         = (pongWait * 9) / 10
	maxMessageSize int64 = 1024
)

type WsConnection struct {
	Cid       string
	Conn      *websocket.Conn
	manager   *WsManager
	ReadChan  chan *MsgPack
	WriteChan chan []byte
	Session   *Session
}

func (c *WsConnection) Run() {
	go c.readMessage()
	go c.writeMessage()
	//做一些心跳检测 websocket中 ping pong机制
	c.Conn.SetPongHandler(c.PongHandler)
}

func (c *WsConnection) Close() {
	if c.Conn != nil {
		_ = c.Conn.Close()
	}
}

func (c *WsConnection) SendMessage(buf []byte) error {
	c.WriteChan <- buf
	return nil
}

func (c *WsConnection) GetSession() *Session {
	return c.Session
}

func (c *WsConnection) readMessage() {
	defer func() {
		c.manager.removeClient(c)
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		global.Logger["err"].Errorf("SetReadDeadline err:%v", err)
		return
	}
	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		//客户端发来的消息处理
		if messageType == websocket.TextMessage {
			if c.ReadChan != nil {
				c.ReadChan <- &MsgPack{
					Cid:  c.Cid,
					Body: message,
				}
			}
		} else {
			global.Logger["err"].Errorf("unsupported message type : %d", messageType)
		}
	}
}

func (c *WsConnection) writeMessage() {
	ticker := time.NewTicker(pingInterval)

	defer ticker.Stop()

	for {
		select {
		case message, ok := <-c.WriteChan:
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					global.Logger["err"].Errorf("connection closed, %v", err)
				}
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				global.Logger["err"].Errorf("client[%s] write message err :%v", c.Cid, err)
			}
		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				global.Logger["err"].Errorf("client[%s] ping SetWriteDeadline err :%v", c.Cid, err)
			}
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				global.Logger["err"].Errorf("client[%s] ping  err :%v", c.Cid, err)
				c.Close()
				return
			}
		}
	}
}

func (c *WsConnection) PongHandler(data string) error {
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return err
	}
	return nil
}

func NewWsConnection(conn *websocket.Conn, manager *WsManager) *WsConnection {
	cid := uuid.New().String()
	return &WsConnection{
		Conn:      conn,
		manager:   manager,
		Cid:       cid,
		WriteChan: make(chan []byte, 1024),
		ReadChan:  manager.ClientReadChan,
		Session:   NewSession(cid),
	}
}
