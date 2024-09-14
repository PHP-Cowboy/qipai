package net

type Connection interface {
	Close()
	SendMessage(buf []byte) error
	GetSession() *Session
}

type MsgPack struct {
	Cid     string `json:"cid"`
	Handler string `json:"handler"`
	Body    []byte `json:"body"`
}
