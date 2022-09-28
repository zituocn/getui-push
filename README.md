# getui push api

* golang版本的个推API v2，仅限ymzy使用；
* 其他用户需要使用，请根据自己情况，修改 `getPushMessageAndChannel`和`getIntent`函数中的配置；



```sh
更多方法请查看 getui.go
```


### init pushClient


```go

var (
	pushClient *push.PushClient
)

func InitClient() {
	conf := &push.PushConfig{
		AppId:        "xxxx",
		AppSecret:    "xxx",
		AppKey:       "xxxx",
		MasterSecret: "xxxx",
	}
	store := &push.PushStore{
		Host:     "127.0.0.1",
		Port:     6379,
		DB:       0,
		Password: "xxxx",
		Key:      "getui:token",
	}
	var err error

    // true 为是否打开调试模式
	pushClient, err = push.NewPushClient(conf, store, true)
	if err != nil {
		logy.Infof("init client error :%v", err)
	}
}

```

### 一些方法

```go

// 获取token
func (g *PushClient) GetToken() (token string, err error) 


// 绑定别名
func (g *PushClient) BindAlias(param *models.Alias) (resp *models.Response, err error) 


// 绑定别名
func (g *PushClient) BindAlias(param *models.Alias) (resp *models.Response, err error) 


// 解绑别名
func (g *PushClient) UnBindAlias(param *models.Alias) (resp *models.Response, err error) 

// 根据cid绑定标签
func (g *PushClient) BindTags(cid string, param *models.CustomTagsParam) (resp *models.Response, err error) 


// 根据cid查询已经绑定的标签
func (g *PushClient) SearchTags(cid string) (resp *models.Response, err error) 


// 根据cid查询用户状态
func (g *PushClient) SearchStatus(cid string) (resp *models.Response, err error) 

// 根据cid查询个推的用户信息
func (g *PushClient) SearchUser(cid string) (resp *models.Response, err error) 

// 按cid查询绑定的别名
func (g *PushClient) SearchAliasByCid(cid string) (resp *models.Response, err error) 

// 按别名查询登录过的设备(cid)
func (g *PushClient) SearchCidByAlias(alias string) (resp *models.Response, err error) 

```

### 推送的方法

```go

// 推给所有人
func (g *PushClient) PushAll(scheduleTime int, payload *models.CustomMessage) (resp *models.Response, err error) 

// 推给指定的客户端:android or iOS
func (g *PushClient) PushAllByClient(scheduleTime int, clientType ClientType, payload *models.CustomMessage) (resp *models.Response, err error) 


// 单推给某一个用户
//  根据cid
func (g *PushClient) PushSingleByCid(channelType int, cid string, payload *models.CustomMessage) (resp *models.Response, err error) 


// 单推给某一个用户
//  根据别名
func (g *PushClient) PushSingleByAlias(channelType int, alias string, payload *models.CustomMessage) (resp *models.Response, err error) 


// 按cid数组进行群推
func (g *PushClient) PushListByCid(cid []string, payload *models.CustomMessage) (data []*models.Response, err error) 

// 按自定义标签进行群推
func (g *PushClient) PushAllByCustomTag(scheduleTime int, customTag []string, payload *models.CustomMessage) (resp *models.Response, err error) 
```

