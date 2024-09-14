package net

import (
	"common/logs"
	"encoding/json"
	"errors"
	"fmt"
	"frame/game"
	"frame/remote"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"sync"
	"time"
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
	handlers           map[string]EventHandler
}

func NewWsManager() *WsManager {
	return &WsManager{
		websocketUpgrade: &websocketUpgrade,
		clients:          make(map[string]Connection),
		ClientReadChan:   make(chan *MsgPack, 1024),
		RemoteReadChan:   make(chan []byte, 1024),
		handlers:         make(map[string]EventHandler),
	}
}

func (m *WsManager) Run(addr string) {
	go m.clientReadChanHandler()
	go m.remoteReadChanHandler()
	http.HandleFunc("/", m.serveWS)
	//设置不同的消息处理器
	m.setupEventHandlers()
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
				m.decodeClientPack(body)

				data, err := json.Marshal(body)
				if err != nil {
					return
				}
				fmt.Println(string(data))
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

func (m *WsManager) decodeClientPack(body *MsgPack) {
	if err := m.routeEvent(body); err != nil {
		logs.Error("routeEvent err:%v", err)
	}
}

func (m *WsManager) routeEvent(body *MsgPack) interface{} {

	//  处理器
	conn, ok := m.clients[body.Cid]
	if ok {
		handler, ok := m.handlers[body.Handler]
		if ok {
			return handler(body, conn)
		} else {
			return errors.New("no Handler found")
		}
	}
	return errors.New("no client found")
}

func (m *WsManager) setupEventHandlers() {
	m.handlers[""] = m.MessageHandler
}

func (m *WsManager) MessageHandler(msg *MsgPack, c Connection) error {
	msg.Handler = "entryHandler.entry"
	//本地connector服务器处理
	handler, ok := m.ConnectorHandlers[msg.Handler]
	if ok {
		data, err := handler(c.GetSession(), msg.Body)
		if err != nil {
			return err
		}

		res, err := json.Marshal(data)
		if err != nil {
			return err
		}

		return c.SendMessage(res)
	}

	return nil
}

func (m *WsManager) selectDst(serverType string) (string, error) {
	serversConfigs, ok := game.Conf.ServersConf.TypeServer[serverType]
	if !ok {
		return "", errors.New("no server found")
	}
	//随机一个 比较好的一个策略
	rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rand.Intn(len(serversConfigs))
	return serversConfigs[index].ID, nil
}
