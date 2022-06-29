package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bigashes/dada_delivery/util"
	"net/http"
	"strconv"
	"time"
)

//生产
const BaseUrl = "http://newopen.imdada.cn"

//测试
const DebugUrl = "http://newopen.qa.imdada.cn"

//请求体
type ReqBody struct {
	Body      string `json:"body"`      //消息体
	Format    string `json:"format"`    //格式
	Timestamp int64  `json:"timestamp"` //时间戳
	Signature string `json:"signature"` //签名
	AppKey    string `json:"app_key"`   //appkey
	V         string `json:"v"`         //版本 1.0
	SourceId  string `json:"source_id"` //商户id
}

//结果体
type RespBody struct {
	Status    string       `json:"status"`
	Code      int64        `json:"code"`
	Msg       string       `json:"msg"`
	Result    *interface{} `json:"result"`
	ErrorCode int64        `json:"errorCode"`
}

var defaultClient = &http.Client{
	Timeout: 60 * time.Second,
}

func NewReqBody(appKey, sourceId, body string) *ReqBody {
	return &ReqBody{
		Body:      body,
		Format:    "json",
		Timestamp: 0,
		Signature: "",
		AppKey:    appKey,
		V:         "1.0",
		SourceId:  sourceId,
	}
}

func PostRequest(incompleteURL string, param *ReqBody, appSecret string, response interface{}, debug bool) error {
	var buf bytes.Buffer
	if param != nil {
		p, err := urlValues(param, appSecret)
		if err != nil {
			return err
		}
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		if err := encoder.Encode(p); err != nil {
			return err
		}
	}

	if debug {
		incompleteURL = DebugUrl + incompleteURL
	} else {
		incompleteURL = BaseUrl + incompleteURL
	}

	httpResp, err := defaultClient.Post(incompleteURL, "application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}

	respBody := &RespBody{
		Result: &response,
	}

	err = json.NewDecoder(httpResp.Body).Decode(respBody)
	if err != nil {
		return err
	}

	if respBody.Status != "success" {
		return util.DecodeErr(respBody.ErrorCode)
	}

	return nil
}

func urlValues(param *ReqBody, appSecret string) (map[string]string, error) {
	value := make(map[string]string)
	value["app_key"] = param.AppKey
	value["timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	value["format"] = param.Format
	value["v"] = param.V
	value["source_id"] = param.SourceId
	value["body"] = param.Body
	value["signature"] = util.Sign(value, appSecret)
	return value, nil
}
