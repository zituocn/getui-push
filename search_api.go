package getuipush

import (
	"encoding/json"
	"github.com/zituocn/getui-push/models"
)

// searchTaskDetailByCid 可以查询某任务下某cid的具体实时推送路径情况
//	此接口需要SVIP权限，暂时不可用
func searchTaskDetailByCid(appId, token string, cid, taskId string) (*models.TaskDetailResp, error) {
	b, err := HttpRequest("GET", appId+"/task/detail/"+cid+"/"+taskId, token, nil)
	if err != nil {
		return nil, err
	}
	resp := new(models.TaskDetailResp)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// searchSchedule 查询定时任务
//	该接口支持在推送完定时任务之后，查看定时任务状态，定时任务是否发送成功。
func searchSchedule(appId, token, taskId string) (*models.Response, error) {
	resp, err := RequestAPI("GET", appId+"/task/schedule/"+taskId, token, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
