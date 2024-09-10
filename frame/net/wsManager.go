package net

import (
	"common/logs"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var (
	websocketUpgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type CheckOriginHandler func(r *http.Request) bool

type WsManager struct {
	sync.RWMutex
	websocketUpgrade   *websocket.Upgrader
	ServerId           string
	CheckOriginHandler CheckOriginHandler
	clients            map[string]Connection
	ConnectorHandlers  LogicHandler
	ClientReadChan     chan *MsgPack
}

func NewWsManager() *WsManager {
	return &WsManager{
		websocketUpgrade: &websocketUpgrade,
		clients:          make(map[string]Connection),
		ClientReadChan:   make(chan *MsgPack, 1024),
	}
}

func (m *WsManager) Run() {

}

func (m *WsManager) Close() {

}

func (m *WsManager) serveWS(w http.ResponseWriter, r *http.Request) {

	wsConn, err := m.websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		logs.Error("websocketUpgrade.Upgrade err:%v", err)
		return
	}
	client := NewWsConnection(wsConn, m)
	m.addClient(client)
	client.Run()
}

func (m *WsManager) addClient(client *WsConnection) {
	m.Lock()
	defer m.Unlock()
	m.clients[client.Cid] = client
}
