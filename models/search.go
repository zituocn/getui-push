package models

// TaskDetailResp 调用此接口可以查询某任务下某cid的具体实时推送路径情况
// {
//     "code":0,
//     "msg":"success",
//     "data":{
//         "deatil":[
//             {
//                 "time":"yyyy-MM-dd HH:mm:ss",
//                 "event":"消息请求成功"
//             },
//             {
//                 "time":"yyyy-MM-dd HH:mm:ss",
//                 "event":"到达客户端"
//             }
//         ]
//     }
// }
type TaskDetailResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Deatil []struct {
			Time  string `json:"time"`
			Event string `json:"event"`
		} `json:"deatil"`
	} `json:"data"`
}
