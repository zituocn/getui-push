package models

// AliasParam 别名绑定参数
type AliasParam struct {
	DataList []*Alias `json:"data_list"`
}

// Alias 别名
type Alias struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}

// Response 统一的返回值
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"` //此处把服务端返回的data处理成了json字串
}

// CustomTagsParam 一个用户绑定多个标签的参数
type CustomTagsParam struct {
	CustomTag []string `json:"custom_tag"`
}
