package user

import (
	"database/sql"
)

type User struct {
	Uid                int64           `db:"uid"`
	Token              sql.NullString  `db:"token"`
	Playername         sql.NullString  `db:"playername"`
	Username           sql.NullString  `db:"username"`
	Phone              sql.NullString  `db:"phone"`
	Isbindfb           sql.NullInt64   `db:"isbindfb"`
	Isbindgg           sql.NullInt64   `db:"isbindgg"`
	Firstbuylist       sql.NullString  `db:"firstbuylist"`
	Agent              sql.NullInt64   `db:"agent"`
	Usericon           sql.NullString  `db:"usericon"`
	Sex                sql.NullInt64   `db:"sex"`
	CreateTime         sql.NullInt64   `db:"create_time"`
	Ddid               sql.NullString  `db:"ddid"`
	CreatePlatform     sql.NullInt64   `db:"create_platform"`
	Coin               sql.NullFloat64 `db:"coin"`
	Dcoin              sql.NullFloat64 `db:"dcoin"`
	Dcashbonus         sql.NullFloat64 `db:"dcashbonus"`
	Cashbonus          sql.NullFloat64 `db:"cashbonus"`
	Totalbet           sql.NullFloat64 `db:"totalbet"`
	Totalwin           sql.NullFloat64 `db:"totalwin"`
	Totaldraw          sql.NullFloat64 `db:"totaldraw"`
	Totalrecharge      sql.NullFloat64 `db:"totalrecharge"`
	Maxdraw            sql.NullFloat64 `db:"maxdraw"`
	Candraw            sql.NullFloat64 `db:"candraw"`
	Gamedraw           sql.NullFloat64 `db:"gamedraw"`
	Gamebonus          sql.NullFloat64 `db:"gamebonus"`
	Points             sql.NullInt64   `db:"points"`
	Charm              sql.NullInt64   `db:"charm"`
	Diamond            sql.NullInt64   `db:"diamond"`
	Rp                 sql.NullInt64   `db:"rp"`
	Code               sql.NullString  `db:"code"`
	Forbidcode         sql.NullInt64   `db:"forbidcode"`
	FromChannel        sql.NullInt64   `db:"from_channel"`
	Memo               sql.NullString  `db:"memo"`
	Adchannel          sql.NullString  `db:"adchannel"`
	LoginTime          sql.NullInt64   `db:"login_time"`
	RegIp              sql.NullString  `db:"reg_ip"`
	LoginIp            sql.NullString  `db:"login_ip"`
	InvitUid           sql.NullInt64   `db:"invit_uid"`
	Status             sql.NullInt64   `db:"status"`
	Sharetime          sql.NullInt64   `db:"sharetime"`
	Ispayer            sql.NullInt64   `db:"ispayer"`
	Totalpay           sql.NullFloat64 `db:"totalpay"`
	Active             sql.NullInt64   `db:"active"`
	LoginDays          sql.NullInt64   `db:"login_days"`
	Wintimes           sql.NullInt64   `db:"wintimes"`
	Alltimes           sql.NullInt64   `db:"alltimes"`
	Isrobot            sql.NullInt64   `db:"isrobot"`
	Iswhite            sql.NullInt64   `db:"iswhite"`
	Isblack            sql.NullInt64   `db:"isblack"`
	Reservecoin        sql.NullInt64   `db:"reservecoin"`
	Limitcoin          sql.NullInt64   `db:"limitcoin"`
	Platform           sql.NullInt64   `db:"platform"`
	Daytime            sql.NullInt64   `db:"daytime"`
	Onlinestate        sql.NullInt64   `db:"onlinestate"`
	Curonlinejd        sql.NullInt64   `db:"curonlinejd"`
	Lrwardstate        sql.NullInt64   `db:"lrwardstate"`
	Praisetime         sql.NullInt64   `db:"praisetime"`
	Invitednum         sql.NullInt64   `db:"invitednum"`
	Invitedfb          sql.NullInt64   `db:"invitedfb"`
	Uuid               sql.NullInt64   `db:"uuid"`
	UpdateTime         sql.NullInt64   `db:"update_time"`
	Spread             sql.NullInt64   `db:"spread"`
	Guide              sql.NullString  `db:"guide"`
	RedEnvelope        sql.NullFloat64 `db:"red_envelope"`
	Hadshowredenvelope sql.NullInt64   `db:"hadshowredenvelope"`
	Svip               sql.NullInt64   `db:"svip"`
	Svipexp            sql.NullInt64   `db:"svipexp"`
	Tagid              sql.NullInt64   `db:"tagid"`
	Level              sql.NullInt64   `db:"level"`
	Levelexp           sql.NullInt64   `db:"levelexp"`
	Ticket             sql.NullInt64   `db:"ticket"`
	DeviceToken        sql.NullString  `db:"deviceToken"`
	Wincoin            sql.NullInt64   `db:"wincoin"`
	Consvalue          sql.NullInt64   `db:"consvalue"`
	Isgetinitgift      sql.NullInt64   `db:"isgetinitgift"`
	Isonetime          sql.NullInt64   `db:"isonetime"`
	Moneybag           sql.NullInt64   `db:"moneybag"`
	MoneybagTime       sql.NullInt64   `db:"moneybag_time"`
	Fcmtoken           sql.NullString  `db:"fcmtoken"`
	Lastgameid         sql.NullInt64   `db:"lastgameid"`
	Ginlamicount       sql.NullInt64   `db:"ginlamicount"`
	Appid              sql.NullInt64   `db:"appid"`
	Justreg            sql.NullInt64   `db:"justreg"`
	SysMailID          sql.NullInt64   `db:"sysMailID"`
	Kouuid             sql.NullString  `db:"kouuid"`
	Country            sql.NullInt64   `db:"country"`
	Avatarframe        sql.NullString  `db:"avatarframe"`
	SysMsgID           sql.NullInt64   `db:"sysMsgID"`
	Identity           sql.NullInt64   `db:"identity"`
	Leagueexp          sql.NullInt64   `db:"leagueexp"`
	Leaguelevel        sql.NullInt64   `db:"leaguelevel"`
	Skinlist           sql.NullString  `db:"skinlist"`
	Chatskin           sql.NullString  `db:"chatskin"`
	Tableskin          sql.NullString  `db:"tableskin"`
	Pokerskin          sql.NullString  `db:"pokerskin"`
	Frontskin          sql.NullString  `db:"frontskin"`
	Emojiskin          sql.NullString  `db:"emojiskin"`
	Faceskin           sql.NullString  `db:"faceskin"`
	Fbicon             sql.NullString  `db:"fbicon"`
	Badge              sql.NullInt64   `db:"badge"`
	Vipendtime         sql.NullInt64   `db:"vipendtime"`
	Verifyfriend       sql.NullInt64   `db:"verifyfriend"`
	Lang               sql.NullInt64   `db:"lang"`
	Fbtoken            sql.NullString  `db:"fbtoken"`
	Fbendtime          sql.NullInt64   `db:"fbendtime"`
	Fgamelist          sql.NullString  `db:"fgamelist"`
	Charmlist          sql.NullString  `db:"charmlist"`
	Blockuids          sql.NullString  `db:"blockuids"`
	Isbindapple        sql.NullInt64   `db:"isbindapple"`
	Salonskin          sql.NullString  `db:"salonskin"`
	Salontesttime      sql.NullInt64   `db:"salontesttime"`
	Countrycn          sql.NullString  `db:"countrycn"`
	Bindbank           sql.NullInt64   `db:"bindbank"`
	Bindupi            sql.NullInt64   `db:"bindupi"`
	Bindusdt           sql.NullInt64   `db:"bindusdt"`
	Isbindphone        sql.NullInt64   `db:"isbindphone"`
	Suspendagent       sql.NullInt64   `db:"suspendagent"`
	Ourselves          sql.NullInt64   `db:"ourself"`
	Forbidchat         sql.NullInt64   `db:"forbidchat"`
	Forbidfriend       sql.NullInt64   `db:"forbidfriend"`
	KYC                sql.NullInt64   `db:"kyc"`
	KYCField           sql.NullInt64   `db:"kycfield"`
	Stopregrebat       sql.NullInt64   `db:"stopregrebat"`
	Forbidnick         sql.NullInt64   `db:"forbidnick"`
	Nevershowtips      sql.NullInt64   `db:"nevershowtips"`
	Payratetype        sql.NullInt64   `db:"payratetype"`
	Drawsucctimes      sql.NullInt64   `db:"drawsucctimes"`
	Drawsucccoin       sql.NullInt64   `db:"drawsucccoin"`
	Device             sql.NullString  `db:"device"`
	Remark             sql.NullString  `db:"remark"`
}
