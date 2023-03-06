package getuipush

const (
	//APIURL 服务器地址
	APIURL string = "https://restapi.getui.com/v2/"

	// NAME 日志中的前缀
	NAME = "[个推]"

	// PublicChannel 运营推送
	PublicChannel = 1 //公信通道

	// PrivateChannel 聊天推送
	PrivateChannel = 2

	// limit 多个cid群推时，每次的用户量
	limit = 1000
)

type XiaoMiMsgType int

const (
	ArticleMsg XiaoMiMsgType = iota + 1
	AlgorithmReCommendMsg
	AttendRecommendMsg
	PlatformActionMsg
	UserAccountMsg
	InstantMsg
)

var xiaomiMagTypeChannelId = map[XiaoMiMsgType]string{
	ArticleMsg:            "103533",
	AlgorithmReCommendMsg: "103779",
	AttendRecommendMsg:    "103777",
	PlatformActionMsg:     "103781",
	UserAccountMsg:        "103776",
	InstantMsg:            "103782",
}

func (m XiaoMiMsgType) GetMsgChannelId() string {
	return xiaomiMagTypeChannelId[m]
}

// ClientType APP客户端类型
type ClientType int

const (
	Android   ClientType = iota + 1 //android
	IOS                             //ios
	WechatAPP                       //微信小程序
)
