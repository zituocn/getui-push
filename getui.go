/*
个推client封装
sam
2022-09-01


pushClient:=NewPushClient(....)
resp,err:=pushClient.PushAll(...)

*/

package getuipush

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zituocn/getui-push/models"
	"strings"
	"time"

	"github.com/zituocn/gow/lib/goredis"
	"github.com/zituocn/gow/lib/logy"
)

var (
	// TTL 消息存放时间
	TTL = 86400000 // 1天： 1 * 24 * 3600 * 1000

	ctx = context.Background()

	// expTime token 在redis中的过期时间
	expTime = time.Now().Add(time.Hour * 20).Unix()
)

// PushConfig 配置
//	从个推获取
type PushConfig struct {
	AppId        string
	AppKey       string
	AppSecret    string
	MasterSecret string
}

// PushStore token存储配置
//	redis配置信息
type PushStore struct {
	Host     string //redis host
	Port     int    // redis port
	DB       int    // redis db
	Password string //redis password
	Key      string //存储key名称
}

// PushClient 个推 push client
type PushClient struct {
	*PushConfig
	*PushStore
}

// NewPushClient 返回个推实例并初始化redis信息
func NewPushClient(conf *PushConfig, store *PushStore, toDebug bool) (client *PushClient, err error) {
	if conf == nil {
		err = errors.New("配置为空")
		return
	}
	if store == nil {
		err = errors.New("存储存储配置为空")
		return
	}
	if conf.AppId == "" || conf.AppSecret == "" || conf.AppKey == "" {
		err = errors.New("个推参数配置不完整")
		return
	}
	if store.Host == "" || store.Port == 0 || store.DB < 0 {
		err = errors.New("存储参数配置不完整")
		return
	}
	if toDebug {
		ToDebug = true
	}
	client = &PushClient{
		PushConfig: conf,
		PushStore:  store,
	}
	err = goredis.InitDefaultDB(&goredis.RedisConfig{
		Host:     store.Host,
		Port:     store.Port,
		Pool:     100,
		Password: store.Password,
		Name:     "gt",
		DB:       store.DB,
	})
	if err != nil {
		return
	}
	return
}

// GetToken 获取token
//	从redis中或api中获取
func (g *PushClient) GetToken() (token string, err error) {
	rdb := goredis.GetRDB()
	token, err = rdb.Get(ctx, g.Key).Result()
	if err != nil {
		logy.Errorf("%s 在redis中获取token失败 :%s", NAME, err.Error())
	}
	if token == "" {
		token, err = getToken(g.AppId, g.AppKey, g.MasterSecret)
		if err != nil {
			err = fmt.Errorf("%s 从API获取token失败: %s", NAME, err.Error())
			return
		}
		_, err = rdb.SetEX(ctx, g.Key, token, time.Duration(expTime)).Result()
		if err != nil {
			logy.Errorf("%s 在redis中存储token失败 :%s", NAME, err.Error())
		}
	}
	return
}

/*
===============================================================
绑定用户别名
===============================================================
*/

// BindAlias 绑定别名
func (g *PushClient) BindAlias(param *models.Alias) (resp *models.Response, err error) {
	if param == nil || param.Cid == "" {
		err = errors.New("param未设置或cid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	dataList := make([]*models.Alias, 0)
	dataList = append(dataList, param)
	aliasParam := &models.AliasParam{
		DataList: dataList,
	}
	return bindAlias(g.AppId, token, aliasParam)
}

// UnBindAlias 解绑别名
//	cid与alias成对出现
func (g *PushClient) UnBindAlias(param *models.Alias) (resp *models.Response, err error) {
	if param == nil || param.Cid == "" {
		err = errors.New("param未设置或cid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	dataList := make([]*models.Alias, 0)
	dataList = append(dataList, param)
	aliasParam := &models.AliasParam{
		DataList: dataList,
	}
	return unBindAlias(g.AppId, token, aliasParam)
}

// UnBindAllAlias 解绑所有与该别名绑定的cid
func (g *PushClient) UnBindAllAlias(alias string) (resp *models.Response, err error) {
	if alias == "" {
		err = errors.New("alias为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return unBindAllAlias(g.AppId, token, alias)
}

/*
===============================================================
绑定自定义标签
===============================================================
*/

// BindTags 一个用户绑定一批标签
//	cid表示用户
func (g *PushClient) BindTags(cid string, param *models.CustomTagsParam) (resp *models.Response, err error) {
	if cid == "" {
		err = errors.New("cid为空")
		return
	}
	if param == nil {
		err = errors.New("param为空")
		return
	}
	if len(param.CustomTag) == 0 {
		err = errors.New("自定义标签长度为0")
		return
	}
	if len(param.CustomTag) > 100 {
		err = errors.New("自定义标签长度大于100个")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return bindTags(g.AppId, token, cid, param)
}

/*
===============================================================
查询相关接口
使用Response.data的返回json，需要进一步格式化展示
===============================================================
*/

// SearchTags 查询某个用户已经绑定的标签
func (g *PushClient) SearchTags(cid string) (resp *models.Response, err error) {
	if cid == "" {
		err = errors.New("cid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return searchTags(g.AppId, token, cid)
}

// SearchStatus 查询某个用户的状态，是否在线，上次在线时间等
//	根据cid查询
func (g *PushClient) SearchStatus(cid string) (resp *models.Response, err error) {
	if cid == "" {
		err = errors.New("cid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return searchStatus(g.AppId, token, cid)
}

// SearchUser 查询用户信息
//	根据cid查询
func (g *PushClient) SearchUser(cid string) (resp *models.Response, err error) {
	if cid == "" {
		err = errors.New("cid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	resp, err = searchUser(g.AppId, token, cid)
	if err != nil {
		return
	}
	return
}

// SearchAliasByCid 按cid查别名
//	即这台设备上登录过哪些帐号
func (g *PushClient) SearchAliasByCid(cid string) (resp *models.Response, err error) {
	if cid == "" {
		err = errors.New("cid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return searchAliasByCid(g.AppId, token, cid)
}

// SearchCidByAlias 按alias查cid
//	即这个alias绑定过哪些设备
func (g *PushClient) SearchCidByAlias(alias string) (resp *models.Response, err error) {
	if alias == "" {
		err = errors.New("别名为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return searchCidByAlias(g.AppId, token, alias)
}

// SearchTaskDetailByCid 可以查询某任务下某cid的具体实时推送路径情况
//	用于跟踪某个用户的消息到达情况
//	此接口需要SVIP权限，暂时不可用
func (g *PushClient) SearchTaskDetailByCid(cid, taskId string) (resp *models.TaskDetailResp, err error) {
	if cid == "" {
		err = errors.New("cid为空")
		return
	}
	if taskId == "" {
		err = errors.New("taskid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return searchTaskDetailByCid(g.AppId, token, cid, taskId)
}

/*
===============================================================
推给所有人
===============================================================
*/

// PushAll 推送给所有人
//	scheduleTime 定时推送时间戳，为0时，不定时
func (g *PushClient) PushAll(scheduleTime int, payload *models.CustomMessage) (resp *models.Response, err error) {
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(PublicChannel, scheduleTime, payload)
	if err != nil {
		return
	}
	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		Audience:    "all",
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}

	resp, err = pushApp(g.AppId, token, pushParam)
	if err != nil {
		return
	}
	return
}

/*
===============================================================
推给指定端类型
===============================================================
*/

// PushAllByClient 推送给不同的客户端
//	clientType 客户端类型，只能选1种
//	scheduleTime 定时推送时间戳，为0时，不定时
func (g *PushClient) PushAllByClient(scheduleTime int, clientType ClientType, payload *models.CustomMessage) (resp *models.Response, err error) {
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(PublicChannel, scheduleTime, payload)
	if err != nil {
		return
	}
	var phones []string
	switch clientType {
	case Android:
		phones = []string{"android"}
	case IOS:
		phones = []string{"ios"}
	case WechatAPP:
		//TODO:
	}

	tag := make([]*models.Tag, 0)
	tag = append(tag, &models.Tag{
		Key:     "phone_type",
		Values:  phones,
		OptType: "or",
	})

	audience := struct {
		Tag []*models.Tag `json:"tag"`
	}{}

	audience.Tag = tag

	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		Audience:    audience,
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}
	resp, err = pushAppByClient(g.AppId, token, pushParam)
	if err != nil {
		return
	}
	return
}

/*
===============================================================
推给某一个用户
===============================================================
*/

// PushSingleByCid 单推给某一个用户
//	cid = 用户的cid信息
//	channelType = 通道类型
func (g *PushClient) PushSingleByCid(channelType int, cid string, payload *models.CustomMessage) (resp *models.Response, err error) {
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(channelType, 0, payload)
	if err != nil {
		return
	}
	audience := struct {
		Cid []string `json:"cid"`
	}{}
	audience.Cid = []string{cid}
	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		Audience:    audience,
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}
	resp, err = pushSingleByCid(g.AppId, token, pushParam)
	if err != nil {
		return
	}
	return
}

/*
===============================================================
推给某一个用户
===============================================================
*/

// PushSingleByAlias 单推给某一个用户
//	alias = 用户的alias
//	channelType = 通道类型
func (g *PushClient) PushSingleByAlias(channelType int, alias string, payload *models.CustomMessage) (resp *models.Response, err error) {
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(channelType, 0, payload)
	if err != nil {
		return
	}

	audience := struct {
		Alias []string `json:"alias"`
	}{}
	audience.Alias = []string{alias}
	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		Audience:    audience,
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}
	resp, err = pushSingleByAlias(g.AppId, token, pushParam)
	if err != nil {
		return
	}
	return
}

/*
===============================================================
按cid群推
===============================================================
*/

//PushListByCid 按cid群推消息
func (g *PushClient) PushListByCid(cid []string, payload *models.CustomMessage) (data []*models.Response, err error) {
	if len(cid) == 0 {
		err = errors.New("cid长度为0")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(PublicChannel, 0, payload)
	if err != nil {
		return
	}

	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}

	// 创建消息
	resp, err := createPushMessage(g.AppId, token, pushParam)
	if err != nil {
		err = fmt.Errorf("%s 保存消息失败: %s", NAME, err.Error())
		return
	}
	//返回的taskId
	taskId := resp.Data
	pageCount := getPageCount(limit, len(cid))
	data = make([]*models.Response, 0)

	// 分页群推
	for i := 1; i <= pageCount; i++ {
		list := getSplitCid(cid, i, limit)
		pushListParam := &models.PushListParam{
			TaskId: taskId,
		}

		pushListParam.Audience.Cid = list //每次的推送列表
		pushListParam.IsAsync = false     //不异步

		respList, err := pushListByCid(g.AppId, token, pushListParam)
		if err != nil {
			logy.Errorf("%s 按cid群推失败: %s %s", NAME, respList.Msg, err.Error())
		}
		data = append(data, respList)
		time.Sleep(time.Microsecond * 500) //休眠500ms
	}

	return
}

/*
===============================================================
根据条件筛选用户推送
===============================================================
*/

// PushAllByCustomTag 对指定应用的符合筛选条件的用户群发推送消息。支持定时、定速功能
//	此接口频次限制100次/天，每分钟不能超过5次(推送限制和接口执行群推共享限制)，定时推送功能需要申请开通才可以使用
//	scheduleTime 定时推送时间戳，为0时，不定时
//	customTag 内的标签是交集的关系
func (g *PushClient) PushAllByCustomTag(scheduleTime int, customTag []string, payload *models.CustomMessage) (resp *models.Response, err error) {
	if len(customTag) == 0 {
		err = errors.New("自定义标签长度为0")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(1, scheduleTime, payload)
	if err != nil {
		return
	}

	tags := make([]*models.Tag, 0)
	tags = append(tags, &models.Tag{
		Key:     "custom_tag",
		Values:  customTag,
		OptType: "or",
	})

	audience := struct {
		Tag []*models.Tag `json:"tag"`
	}{}

	audience.Tag = tags

	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		Audience:    audience,
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}
	resp, err = pushAppByTag(g.AppId, token, pushParam)
	if err != nil {
		return
	}
	return
}

/*
===============================================================
使用标签快速推送
===============================================================
*/

// PushAppByFastCustomTag 使用标签快速推送
//	tag 为某一个标签名
//	scheduleTime 为定时任务的时间戳
//	此接口需要SVIP才有使用权限
func (g *PushClient) PushAppByFastCustomTag(scheduleTime int, tag string, payload *models.CustomMessage) (resp *models.Response, err error) {
	if tag == "" {
		err = errors.New("自定义标签长度为0")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	pushMessage, pushChannel, setting, err := getPushMessageAndChannel(1, scheduleTime, payload)
	if err != nil {
		return
	}

	audience := struct {
		FastCustomTag string `json:"fast_custom_tag"`
	}{}

	audience.FastCustomTag = tag
	pushParam := &models.PushParam{
		GroupName:   getGroupName(),
		RequestId:   getRandString(),
		Setting:     setting,
		Audience:    audience,
		PushMessage: pushMessage,
		PushChannel: pushChannel,
	}
	resp, err = pushAppByFastCustomTag(g.AppId, token, pushParam)
	if err != nil {
		return
	}
	return
}

/*
===============================================================
使用标签快速推送
===============================================================
*/

// StopTask 停止推送任务
//	对正处于推送状态，或者未接收的消息停止下发（只支持批量推和群推任务）
func (g *PushClient) StopTask(taskId string) (resp *models.Response, err error) {
	if taskId == "" {
		err = errors.New("taskid为空")
		return
	}
	token, err := g.GetToken()
	if err != nil {
		return
	}
	return stopTask(g.AppId, token, taskId)
}

/*
private
*/

// getPushMessageAndChannel 构造消息
//	channelType 通道
//	scheduleTime 定时任务的时间戳
//	payload 消息结构体
func getPushMessageAndChannel(channelType int, scheduleTime int, payload *models.CustomMessage) (pushMessage *models.PushMessage, pushChannel *models.PushChannel, setting *models.Setting, err error) {
	payload.Title = strings.TrimSpace(payload.Title)
	pushInfo, err := json.Marshal(payload)
	if err != nil {
		return
	}

	// 参数配置
	setting = &models.Setting{
		TTL: TTL,
	}
	setting.Strategy.IOS = 2
	setting.Strategy.Default = 1
	if scheduleTime > 0 {
		setting.ScheduleTime = scheduleTime
	}

	// 个推消息，走透传模式
	// TODO:此处可测试是否可走 通知消息模式
	pushMessage = &models.PushMessage{
		Transmission: string(pushInfo),
	}

	// iOS消息配置
	ios := &models.IOSChannel{
		Payload:   string(pushInfo),
		Type:      "notify",
		AutoBadge: "+1",
	}
	ios.Aps.ContentAvailable = 0 //通知消息 =1时为静默消息
	ios.Aps.Sound = "default"    //铃声
	ios.Aps.Alert.Title = payload.Title
	ios.Aps.Alert.Body = payload.Content

	// android 消息配置
	android := &models.AndroidChannel{}

	//走厂商的通知消息
	android.Ups.Notification = &models.UPSNotification{
		Title:     payload.Title,
		Body:      payload.Content,
		ClickType: "intent",
		Intent:    getIntent(payload.Url),
		NotifyId:  uint(time.Now().Unix()),
	}

	// android 离线推送通道
	// 以下为厂商配置
	android.Ups.Options.All.Channel = "yuanmeng_push"

	//华为 OK except p10
	android.Ups.Options.Hw = map[string]interface{}{
		"/message/android/notification/channel_id": "yuanmeng_push",
		"/message/android/notification/visibility": "PUBLIC",
		"/message/android/notification/importance": "HIGH",
	}

	// 小米
	android.Ups.Options.Xm = map[string]interface{}{
		"/extra.channel_id": "pre213",
	}

	// oppo
	if channelType == PublicChannel {
		android.Ups.Options.Op.ChannelId = "yuanmeng_push"
	}
	if channelType == PrivateChannel {
		android.Ups.Options.Op.ChannelId = "yuanmeng_push_im"
	}

	//vivo
	android.Ups.Options.Vv.Classification = 1
	android.Ups.Options.Vv.NotifyType = 4

	pushChannel = &models.PushChannel{
		Android: android,
		IOS:     ios,
	}

	return
}

// getIntent 返回android的intent地址
func getIntent(url string) string {
	if url == "" {
		return ""
	}
	intent := fmt.Sprintf("intent:#Intent;launchFlags=0x4000000;component=com.yuanmengzhiyuan.ei8z.yuanmeng_app/.module.appHome.MainActivity;S.nextPage=%s;end", url)
	return intent
}

func getRandString() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

func getGroupName() string {
	return fmt.Sprintf("ymzy_%d", time.Now().Year())
}

// getSplitCid return cut []string
func getSplitCid(cid []string, p, limit int) []string {
	list := make([]string, 0)
	offset := (p - 1) * limit
	count := len(cid)
	for i := offset; i < offset+limit && i < count; i++ {
		list = append(list, cid[i])
	}
	return list
}

// getPageCount return pageCount
func getPageCount(limit, count int) (pageCount int) {
	if count > 0 && limit > 0 {
		if count%limit == 0 {
			pageCount = count / limit
		} else {
			pageCount = count/limit + 1
		}
	}
	return pageCount
}
