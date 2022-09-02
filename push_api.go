package getuipush

import (
	"github.com/tidwall/gjson"
	"github.com/zituocn/getui-push/models"
)

// pushSingleByCid 推送给单个用户
//	cid在param中设置
func pushSingleByCid(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/single/cid", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// pushSingleByAlias 推送给单个用户
//	alias在param中设置
func pushSingleByAlias(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/single/alias", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// pushApp 推给所有
func pushApp(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/all", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// pushAppByClient 推给不同客户端
//	客户端指android或ios
//	是android还是ios，从param中区别
func pushAppByClient(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/tag", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// pushAppByTag 推给不同的tag
//	自定义tag
func pushAppByTag(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/tag", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// pushAppByFastCustomTag 使用标签快速推送
func pushAppByFastCustomTag(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/fast_custom_tag", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// createPushMessage 此接口用来创建消息体，并返回taskid，为批量推的前置步骤
//	taskid 任务编号，用于执行cid批量推和执行别名批量推，此taskid可以多次使用，有效期为离线时间
func createPushMessage(appId, token string, param *models.PushParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	b, err := HttpRequest("POST", appId+"/push/list/message", token, bodyByte)
	if err != nil {
		return nil, err
	}
	resp := &models.Response{
		Code: int(gjson.GetBytes(b, "code").Int()),
		Msg:  gjson.GetBytes(b, "msg").String(),
		Data: gjson.GetBytes(b, "data.taskid").String(),
	}
	return resp, nil
}

// pushListByCid 按cid群推
//	使用前，请先调用 CreatePushMessage 后返回的taskid
func pushListByCid(appId, token string, param *models.PushListParam) (*models.Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/push/list/cid", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// stopTask 停止任务
//	对正处于推送状态，或者未接收的消息停止下发（只支持批量推和群推任务）
func stopTask(appId, token, taskId string) (*models.Response, error) {
	resp, err := RequestAPI("DELETE", appId+"/task/"+taskId, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
