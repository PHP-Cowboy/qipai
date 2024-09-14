package net

type HandlerFunc func(session *Session, body []byte) (any, error)

type LogicHandler map[string]HandlerFunc

type EventHandler func(msg *MsgPack, c Connection) error
