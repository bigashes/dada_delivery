package util

import (
	"github.com/pkg/errors"
	"sort"
	"strings"
)

//请求签名
func Sign(p map[string]string, appSecret string) string {
	// 按key顺序整理出来
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var strParam string
	for _, k := range keys {
		if k != "sign" {
			strParam += k + p[k]
		}
	}
	// 首加上app_secret秘钥
	strParam = appSecret + strParam + appSecret

	return strings.ToUpper(Md5(strParam))
}

//订单回调验证签名
func CallbackSign(params []string, signature string) error {

	sort.Strings(params)
	var strParam string
	for _, item := range params {
		strParam += item
	}

	sign := Md5(strParam)
	if sign != signature {
		return errors.New("签名错误")
	}
	return nil
}
