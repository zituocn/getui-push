package getuipush

// Setting 配置
// @https://docs.getui.com/getui/server/rest_v2/common_args/?id=doc-title-6
// strategy:
// 默认所有通道的策略选择1-4
// 1: 表示该消息在用户在线时推送个推通道，用户离线时推送厂商通道;
// 2: 表示该消息只通过厂商通道策略下发，不考虑用户是否在线;
// 3: 表示该消息只通过个推通道下发，不考虑用户是否在线；
// 4: 表示该消息优先从厂商通道下发，若消息内容在厂商通道代发失败后会从个推通道下发。
type Setting struct {
	TTL      int `json:"ttl"` //消息离线时间设置，单位毫秒，-1表示不设离线，-1 ～ 3 * 24 * 3600 * 1000(3天)之间
	Strategy struct {
		Default int `json:"default"`
		IOS     int `json:"ios"`
		ST      int `json:"st"` //锤子/坚果
		HW      int `json:"hw"` //华为
		XM      int `json:"xm"` //小米
		VV      int `json:"vv"` //vivo
		MZ      int `json:"mz"` //魅族
		OP      int `json:"op"` //oppo
	} `json:"strategy"` //厂商通道策略 1~4
	Speed        int `json:"speed"`         //定速推送，例如100，个推控制下发速度在100条/秒左右，0表示不限速
	ScheduleTime int `json:"schedule_time"` //定时推送时间，必须是7天内的时间，格式：毫秒时间戳
}

// Notification 通知消息
//	在线个推通道消息内容
//	仅支持安卓系统，iOS系统不展示个推通道下发的通知消息
//	@https://docs.getui.com/getui/server/rest_v2/common_args/?id=doc-title-6
type Notification struct {
	Title        string `json:"title"`         //标题
	Body         string `json:"body"`          //内容
	ClickType    string `json:"click_type"`    //intent:打开应用内特定页；url:打开网页；payload:自定义消息内容启动应用;payload_custom:自定义消息不启动应用;startapp:打开应用首页；none:纯通知，无动作；
	Intent       string `json:"intent"`        //client_type为intent时填写；
	Url          string `json:"url"`           //client_type为url时填写
	Payload      string `json:"payload"`       //client_type为payload相关时
	NotifyId     uint   `json:"notify_id"`     //覆盖任务时，两条消息的notify_id相同，会覆盖上一条；
	BadgeAddNum  uint   `json:"badge_add_num"` //必须大于0；举例：角标数字配置1，应用之前角标数为2，发送此角标消息后，应用角标数显示为3
	ChannelId    string `json:"channel_id"`    //通知渠道id
	ChannelName  string `json:"channel_name"`  //通知渠道名称
	ChannelLevel int    `json:"channel_level"` //通知渠道重要性：0 1 2 3 4
}

// UPSNotification android厂商的 notification
//client_type:
// 点击通知后续动作,
// 目前支持以下后续动作，
// intent：打开应用内特定页面(厂商都支持)，
// url：打开网页地址(厂商都支持；华为要求https协议，且游戏类应用不支持打开网页地址)，
// startapp：打开应用首页(厂商都支持)
type UPSNotification struct {
	Title     string `json:"title"`      //标题
	Body      string `json:"body"`       //内容
	ClickType string `json:"click_type"` //
	Url       string `json:"url"`        //client_type为url时填写
	Intent    string `json:"intent"`     //client_type为intent时填写；点击通知打开应用特定页面，intent格式必须正确且不能为空，长度 ≤ 4096;【注意：vivo侧厂商限制 ≤ 1024】
	NotifyId  uint   `json:"notify_id"`  //覆盖任务时，两条消息的notify_id相同，会覆盖上一条；
}

// AndroidChannel android 厂商通道消息
type AndroidChannel struct {
	Ups struct {
		Notification *UPSNotification `json:"notification"` //通知消息内容，与transmission、revoke 三选一，都填写时报错。若希望客户端离线时，直接在系统通知栏中展示通知栏消息，推荐使用此参数。
		Transmission string           `json:"transmission"` //透传消息内容，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
		Options      struct {
			All struct {
				Channel string `json:"channel"`
			} `json:"ALL"`
			Hw map[string]interface{} `json:"HW"`
			Op struct {
				ChannelId string `json:"channel_id"`
			} `json:"OP"`
			Vv struct {
				Classification int `json:"classification"` //  0代表运营消息，1代表系统消息
				NotifyType     int `json:"notifyType"`     //通知类型 1:无，2:响铃，3:振动，4:响铃和振动 注意：只对Android 8.0及以下系统有效
			} `json:"VV"`
			Xm map[string]interface{} `json:"XM"`
		} `json:"options"` //第三方厂商扩展内容
	} `json:"ups"`
}

// IOSChannel ios厂商通道消息
type IOSChannel struct {
	Type    string `json:"type"`    //voip：voip语音推送，notify：apns通知消息；notify默认通知消息
	Payload string `json:"payload"` //自定义消息内容
	Aps     struct {
		Alert struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"alert"`
		ContentAvailable int    `json:"content-available"` //0:表示普通通知 1:表示静默消息
		Sound            string `json:"sound"`             //铃声，默认即可
	} `json:"aps"` //推送通知消息内容
	AutoBadge string `json:"auto_badge"` //用于计算icon上显示的数字，还可以实现显示数字的自动增减，如“+1”、 “-1”、 “1” 等，计算结果将覆盖badge
}

// PushChannel 厂商通道消息
type PushChannel struct {
	IOS     *IOSChannel     `json:"ios"`
	Android *AndroidChannel `json:"android"`
}

type Revoke struct {
	OldTaskId string `json:"old_task_id"`
}

// PushMessage 在线个推通道消息内容
type PushMessage struct {
	Notification *Notification `json:"notification"` //通知消息内容，仅支持安卓系统，iOS系统不展示个推通知消息，与transmission、revoke三选一，都填写时报错
	Transmission string        `json:"transmission"` //纯透传消息内容，安卓和iOS均支持，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
	Revoke       *Revoke       `json:"revoke"`       //撤回消息时使用(仅撤回个推通道消息)，与notification、transmission三选一，都填写时报错(消息撤回请勿填写策略参数)
}

// PushParam 推送上报参数
type PushParam struct {
	RequestId   string       `json:"request_id"`   //请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name"`   //任务组名。多个消息任务可以用同一个任务组名，后续可根据任务组名查询推送情况（长度限制100字符，且不能含有特殊符号）只允许填写数字、字母、横杠、下划线
	Setting     *Setting     `json:"setting"`      //配置
	Audience    interface{}  `json:"audience"`     //推送的目标用户，可能包括：cid,alias,tag,all等，根据具体情况动态;包括android|ios @see http://docs.getui.com/getui/server/rest_v2/common_args/?id=doc-title-3
	PushMessage *PushMessage `json:"push_message"` //个推通道消息内容
	PushChannel *PushChannel `json:"push_channel"` //厂商通道
}

// CustomMessage 自定义的消息处理结构体
type CustomMessage struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	Pic          string `json:"pic"`
	Url          string `json:"url"`
	Time         int    `json:"time"`
	IsShowNotify string `json:"is_show_notify"`
}

// Tag 自定义标签
type Tag struct {
	Key     string   `json:"key"`
	Values  []string `json:"values"`
	OptType string   `json:"opt_type"`
}

// PushListParam 按cid或alias群推时的结构体
type PushListParam struct {
	Audience struct {
		Cid []string `json:"cid"` //cid数组长度不能大于1000
	} `json:"audience"`
	TaskId  string `json:"taskid"`
	IsAsync bool   `json:"is_async"`
}
