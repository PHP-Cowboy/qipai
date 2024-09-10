package remote

type Msg struct {
	Cid         string
	Body        []byte
	Src         string
	Dst         string
	Router      string
	Uid         string
	SessionData map[string]any
	Type        int // 0 normal 1 session
	PushUser    []string
}

const SessionType = 1
