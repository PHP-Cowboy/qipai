package net

import (
	"common/logs"
	"fmt"
	"frame/remote"
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
	RemoteReadChan     chan []byte
	RemoteCli          remote.Client
}

func NewWsManager() *WsManager {
	return &WsManager{
		websocketUpgrade: &websocketUpgrade,
		clients:          make(map[string]Connection),
		ClientReadChan:   make(chan *MsgPack, 1024),
		RemoteReadChan:   make(chan []byte, 1024),
	}
}

func (m *WsManager) Run(addr string) {
	go m.clientReadChanHandler()
	go m.remoteReadChanHandler()
	http.HandleFunc("/", m.serveWS)

	logs.Fatal("connector listen serve err:%v", http.ListenAndServe(addr, nil))
}

func (m *WsManager) Close() {
	for cid, v := range m.clients {
		v.Close()
		delete(m.clients, cid)
	}
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

func (m *WsManager) clientReadChanHandler() {
	for {
		select {
		case body, ok := <-m.ClientReadChan:
			if ok {
				fmt.Println(body)
			}
		}
	}
}

func (m *WsManager) remoteReadChanHandler() {
	for {
		select {
		case msg := <-m.RemoteReadChan:
			logs.Info("sub nats msg:%v", string(msg))
		}
	}
}

func (m *WsManager) removeClient(wc *WsConnection) {
	for cid, c := range m.clients {
		if cid == wc.Cid {
			c.Close()
			delete(m.clients, cid)
		}
	}
}
