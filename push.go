package getuipush

import "github.com/tidwall/gjson"

// bindAlias 绑定别名
// @https://docs.getui.com/getui/server/rest_v2/user/
func bindAlias(appId, token string, param *AliasParam) (*Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/user/alias", token, bodyByte)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// unBindAlias 解绑别名
func unBindAlias(appId, token string, param *AliasParam) (*Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("DELETE", appId+"/user/alias", token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// bindTags 给一个cid，绑定多个标签
//	此接口对单个cid有频控限制，每天只能修改一次，最多设置100个标签；单个标签长度最大为32字符，标签总长度最大为512个字符
func bindTags(appId, token, cid string, param *CustomTagsParam) (*Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	resp, err := RequestAPI("POST", appId+"/user/custom_tag/cid/"+cid, token, bodyByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// searchTags 查询某个用户已绑定的标签
//	可用于运营后台查询
func searchTags(appId, token, cid string) (string, error) {
	ret, err := HttpRequest("GET", appId+"/user/custom_tag/cid/"+cid, token, nil)
	if err != nil {
		return "", err
	}
	tags := gjson.GetBytes(ret, "data."+cid)
	return tags.String(), nil
}

// pushSingleByCid 推送给单个用户
//	cid在param中设置
func pushSingleByCid(appId, token string, param *PushParam) (*Response, error) {
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
func pushSingleByAlias(appId, token string, param *PushParam) (*Response, error) {
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
func pushApp(appId, token string, param *PushParam) (*Response, error) {
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
func pushAppByClient(appId, token string, param *PushParam) (*Response, error) {
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
func pushAppByTag(appId, token string, param *PushParam) (*Response, error) {
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
func pushAppByFastCustomTag(appId, token string, param *PushParam) (*Response, error) {
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
func createPushMessage(appId, token string, param *PushParam) (*Response, error) {
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return nil, err
	}
	b, err := HttpRequest("POST", appId+"/push/list/message", token, bodyByte)
	if err != nil {
		return nil, err
	}
	resp := &Response{
		Code: int(gjson.GetBytes(b, "code").Int()),
		Msg:  gjson.GetBytes(b, "msg").String(),
		Data: gjson.GetBytes(b, "data.taskid").String(),
	}
	return resp, nil
}

// pushListByCid 按cid群推
//	使用前，请先调用 CreatePushMessage 后返回的taskid
func pushListByCid(appId, token string, param *PushListParam) (*Response, error) {
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
