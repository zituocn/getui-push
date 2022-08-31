package getuipush

const (
	//APIURL 服务器地址
	APIURL string = "https://resetapi.getui.com/v2/"

	// NAME 日志中的前缀
	NAME = "[个推]"

	// PublicChannel 运营推送
	PublicChannel = 1 //公信通道

	// PrivateChannel 聊天推送
	PrivateChannel = 2

	// limit 多个cid群推时，每次的用户量
	limit = 1000
)

// ClientType APP客户端类型
type ClientType int

const (
	Android   ClientType = iota + 1 //android
	IOS                             //ios
	WechatAPP                       //微信小程序
)
