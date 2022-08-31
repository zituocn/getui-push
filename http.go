package getuipush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

var (
	// ToDebug 全局的调试开关
	ToDebug = false
)

// RequestAPI 请求API，返回Response
func RequestAPI(method, url, token string, bodyByte []byte) (*Response, error) {
	data, err := HttpRequest(method, url, token, bodyByte)
	if err != nil {
		return nil, err
	}
	code := gjson.GetBytes(data, "code")
	if code.Int() != 0 {
		msg := gjson.GetBytes(data, "msg")
		return nil, fmt.Errorf("%s 请求接口 %s 返回错误代码: %s 信息: %s", NAME, method+" "+url, code, msg)
	}
	resp := &Response{
		Code: int(gjson.GetBytes(data, "code").Int()),
		Msg:  gjson.GetBytes(data, "msg").String(),
		Data: gjson.GetBytes(data, "data").String(),
	}
	return resp, nil
}

// HttpRequest 请求API,返回 []byte
func HttpRequest(method, url, token string, bodyByte []byte) ([]byte, error) {
	u := APIURL + url
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	body := bytes.NewBuffer(bodyByte)
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("token", token)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if ToDebug {
		debugPrint("Request Method", method)
		debugPrint("Request URL", url)
		debugPrint("Request Header", req.Header)
		debugPrint("Request Body", string(bodyByte))
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if ToDebug {
		debugPrint("Response Status", fmt.Sprintf("%d", resp.StatusCode))
		debugPrint("Response Header", resp.Header)
		debugPrint("Response Body", string(ret))
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s response error , status: %d body: %s", NAME, resp.StatusCode, string(ret))
		return nil, err
	}
	return ret, nil
}

// makeReqBody 序列号v to json []byte
func makeReqBody(v interface{}) ([]byte, error) {
	if ToDebug {
		body, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return nil, err
		}
		return body, nil
	} else {
		body, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
}

// debugPrint 打印调试信息
func debugPrint(prefix string, v interface{}) {
	fmt.Printf("%s %v \n", leftText(prefix+":"), v)
}

func leftText(s string) string {
	return fmt.Sprintf("%15s", s)
}
