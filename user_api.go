package getuipush

import "github.com/zituocn/getui-push/models"

// bindAlias 绑定别名
// @https://docs.getui.com/getui/server/rest_v2/user/
func bindAlias(appId, token string, param *models.AliasParam) (*models.Response, error) {
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
//	cid与alias成对出现
func unBindAlias(appId, token string, param *models.AliasParam) (*models.Response, error) {
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

// unBindAllAlias 解绑所有与该别名绑定的cid
func unBindAllAlias(appId, token, alias string) (*models.Response, error) {
	resp, err := RequestAPI("DELETE", appId+"/user/alias/"+alias, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// bindTags 给一个cid，绑定多个标签
//	此接口对单个cid有频控限制，每天只能修改一次，最多设置100个标签；单个标签长度最大为32字符，标签总长度最大为512个字符
func bindTags(appId, token, cid string, param *models.CustomTagsParam) (*models.Response, error) {
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
func searchTags(appId, token, cid string) (*models.Response, error) {
	resp, err := RequestAPI("GET", appId+"/user/custom_tag/cid/"+cid, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// searchStatus 查询某个用户的状态，是否在线，上次在线时间等
//	根据cid查询
func searchStatus(appId, token, cid string) (*models.Response, error) {
	resp, err := RequestAPI("GET", appId+"/user/status/"+cid, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// searchUser 查询用户信息
//	根据cid查询
func searchUser(appId, token, cid string) (*models.Response, error) {
	resp, err := RequestAPI("GET", appId+"/user/detail/"+cid, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// searchAliasByCid 按cid查询别名
//	即这台设备上登录过哪些帐号
func searchAliasByCid(appId, token, cid string) (*models.Response, error) {
	resp, err := RequestAPI("GET", appId+"/user/alias/cid/"+cid, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// searchCidByAlias 按alias查cid
//	即这个alias绑定过哪些设备
func searchCidByAlias(appId, token, alias string) (*models.Response, error) {
	resp, err := RequestAPI("GET", appId+"/user/cid/alias/"+alias, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
