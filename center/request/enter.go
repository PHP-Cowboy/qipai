package request

type Enter struct {
}

type T struct {
	C        int    `json:"c"`
	Uid      int    `json:"uid"`
	Openid   string `json:"openid"`
	Server   string `json:"server"`
	Subid    int    `json:"subid"`
	Token    string `json:"token"`
	Deviceid string `json:"deviceid"`
	Appver   string `json:"appver"`
	CTs      int64  `json:"c_ts"`
	CIdx     int    `json:"c_idx"`
	Language int    `json:"language"`
}
