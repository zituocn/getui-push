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

type MessageType int

const (
	ArticleMsg MessageType = iota + 1
	AlgorithmReCommendMsg
	AttendRecommendMsg
	PlatformActionMsg
	UserAccountMsg
	InstantMsg
)

var messageTypeText = map[MessageType]string{
	ArticleMsg:            "内容资讯",
	AlgorithmReCommendMsg: "算法推荐",
	AttendRecommendMsg:    "关注推荐",
	PlatformActionMsg:     "平台活动",
	UserAccountMsg:        "个人账户",
	InstantMsg:            "即时消息/聊天消息",
}

// 小米 channelId 定义在小米开发者后台
var xiaomiMessageTypeChannelId = map[MessageType]string{
	ArticleMsg:            "103533",
	AlgorithmReCommendMsg: "103779",
	AttendRecommendMsg:    "103777",
	PlatformActionMsg:     "103781",
	UserAccountMsg:        "103776",
	InstantMsg:            "103782",
}

/*
华为 category意义：

ACCOUNT（帐号动态）：
用户帐号和帐号下资源资产的动态信息。
帐号：帐号上下线、帐号状态变化、帐号信息认证等。
资产：会员到期/过期、续费提醒、余额变动（余额必须为真实的资产变动，且需排除积分变动、金币变动，排名更新等）。

MARKETING（资讯营销类）：
包含 内容资讯 和 营销活动 两大类
*/

// 华为 category 定义在华为开发者后台
var huaweiMessageTypeCategory = map[MessageType]string{
	ArticleMsg:            "MARKETING",
	AlgorithmReCommendMsg: "MARKETING",
	AttendRecommendMsg:    "MARKETING",
	PlatformActionMsg:     "MARKETING", //由于小米的分类比较详细，后台需要分为四类，对华为来说，都归到资讯营销类
	UserAccountMsg:        "ACCOUNT",   //账号动态
	InstantMsg:            "IM",        //即时聊天
}

// 华为channelId 定义在客户端
var huaweiMessageTypeChannelId = map[MessageType]string{
	ArticleMsg:            "yuanmeng_push",
	AlgorithmReCommendMsg: "yuanmeng_push",
	AttendRecommendMsg:    "yuanmeng_push",
	PlatformActionMsg:     "yuanmeng_push",
	UserAccountMsg:        "yuanmeng_push_user",
	InstantMsg:            "yuanmeng_push_im",
}

// 华为 自定义铃声
var huaweiMessageTypeImportance = map[MessageType]string{
	ArticleMsg:            "LOW",
	AlgorithmReCommendMsg: "LOW",
	AttendRecommendMsg:    "LOW",
	PlatformActionMsg:     "LOW",
	UserAccountMsg:        "NORMAL",
	InstantMsg:            "NORMAL",
}

// 荣耀 消息类型: LOW 资讯营销类，NORMAL 服务通讯类
var honorMessageTypeImportance = map[MessageType]string{
	ArticleMsg:            "LOW",
	AlgorithmReCommendMsg: "LOW",
	AttendRecommendMsg:    "LOW",
	PlatformActionMsg:     "LOW",
	UserAccountMsg:        "NORMAL",
	InstantMsg:            "NORMAL",
}

/*
VIVO category意义：

MARKETING（运营活动）：
1. 非用户主动设置，需用户参与的活动提醒、小游戏提醒、服务或商品评价提醒等。 如：抽奖、积分、签到、任务、分享、偷菜、领金币等；
2. 商品推荐，包括红包折扣、商家服务更新、店铺上新等。如：可能感兴趣、商品达到最低价、满减、促销、返利、优惠券、代金券、送红包、信用分增加等相关的通知；
3. 其他消息：用户调研问卷、功能介绍、邀请推荐、版本更新等。

CONTENT（内容推荐）：
内容型的信息推荐，包含热搜、点评、广告、书籍、音乐、视频、直播、课程、节目、游戏宣传、社区话题等。以及：
1. 各垂直类目的相关内容资讯。
2. 天气预报：包括各类天气预报、天气预警提醒等。
3. 出行资讯：包括交规公告、驾考信息、导航路况、铁路购票公告、疫情消息，道路管控等。

ACCOUNT（账号与资产）：
账号变动：帐号上下线、状态变化、信息认证、会员到期、续费提醒、余额变动等。
资产变动：账户下的真实资产变动，交易提示、话费余额、流量、语音时长、短信额度等典型运营商提醒。
*/

// vivo category
var vivoMessageTypeCategory = map[MessageType]string{
	ArticleMsg:            "CONTENT",
	AlgorithmReCommendMsg: "CONTENT",
	AttendRecommendMsg:    "CONTENT",
	PlatformActionMsg:     "MARKETING", //运营活动
	UserAccountMsg:        "ACCOUNT",   //账号与资产
	InstantMsg:            "IM",        //即时消息
}

/*
VIVO classification:
系统消息（1）：即时消息、账号与资产、日程待办、设备信息、订单与物流、订阅提醒
运营消息（0）：新闻、内容推荐、运营活动、社交动态
*/
var vivoMessageTypeClassification = map[MessageType]int64{
	ArticleMsg:            0,
	AlgorithmReCommendMsg: 0,
	AttendRecommendMsg:    0,
	PlatformActionMsg:     0, //运营消息
	UserAccountMsg:        1,
	InstantMsg:            1, //系统消息
}

// oppo channelId
var oppoMessageTypeChannelId = map[MessageType]string{
	ArticleMsg:            "yuanmeng_push",
	AlgorithmReCommendMsg: "yuanmeng_push",
	AttendRecommendMsg:    "yuanmeng_push",
	PlatformActionMsg:     "yuanmeng_push",      //运营消息 公信通道
	UserAccountMsg:        "yuanmeng_push_user", //个人账号 私信通道
	InstantMsg:            "yuanmeng_push_im",   //即时消息 私信通道
}

// GetXiaoMiChannelId 小米: 不同消息类型对应的channelId
func (m MessageType) GetXiaoMiChannelId() string {
	var channelId string = "103533" //默认 内容资讯
	if xiaomiMessageTypeChannelId[m] != "" {
		channelId = xiaomiMessageTypeChannelId[m]
	}
	return channelId
}

// GetHuaweiCategory 华为: 不同消息类型对应的category
func (m MessageType) GetHuaweiCategory() string {
	category := "MARKETING"
	if huaweiMessageTypeCategory[m] != "" {
		category = huaweiMessageTypeCategory[m]
	}
	return category
}

// GetHuaweiChannelId 华为 channelId
func (m MessageType) GetHuaweiChannelId() string {
	channelId := "yuanmeng_push"
	if huaweiMessageTypeChannelId[m] != "" {
		channelId = huaweiMessageTypeChannelId[m]
	}
	return channelId
}

func (m MessageType) GetHuaweiImportance() string {
	return huaweiMessageTypeImportance[m]
}

// GetHuaweiInfo 华为 channelId category
func (m MessageType) GetHuaweiInfo() (channelId, category, importance string) {
	channelId, category, importance = "yuanmeng_push", "MARKETING", "LOW"
	if huaweiMessageTypeChannelId[m] != "" {
		channelId = huaweiMessageTypeChannelId[m]
	}
	if huaweiMessageTypeCategory[m] != "" {
		category = huaweiMessageTypeCategory[m]
	}
	if huaweiMessageTypeImportance[m] != "" {
		importance = huaweiMessageTypeImportance[m]
	}
	return
}

func (m MessageType) GetHonorImportance() (importance string) {
	importance = "LOW" //LOW: 资讯营销类
	if honorMessageTypeImportance[m] != "" {
		importance = honorMessageTypeImportance[m]
	}
	return
}

// GetViVoCategory vivo: 不同消息类型对应的category
func (m MessageType) GetViVoCategory() string {
	category := "CONTENT"
	if vivoMessageTypeCategory[m] != "" {
		category = vivoMessageTypeCategory[m]
	}
	return category
}

// GetViVoClassification vivo: 不同消息类型对应的分类
func (m MessageType) GetViVoClassification() int64 {
	return vivoMessageTypeClassification[m]
}

func (m MessageType) GetOPPOChannelId() string {
	channelId := "yuanmeng_push"
	if oppoMessageTypeChannelId[m] != "" {
		channelId = oppoMessageTypeChannelId[m]
	}
	return channelId
}

// ClientType APP客户端类型
type ClientType int

const (
	Android   ClientType = iota + 1 //android
	IOS                             //ios
	WechatAPP                       //微信小程序
)
