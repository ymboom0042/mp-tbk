/**
 * @Author: YMBoom
 * @Description:
 * @File:  ddx
 * @Version: 1.0.0
 * @Date: 2022/04/18 15:01
 */
package apis

import (
	"encoding/json"
	"github.com/mvdan/xurls"
	"github.com/spf13/viper"
	"mp-17208-top/pkg/resp"
	"mp-17208-top/pkg/utils"
	"net/url"
	"strings"
)

type Ddx struct {
	Gateway, ApiKey string
}

func (d *Ddx) VipItemLink(content string) (tkl string) {
	if content == "" {
		return
	}

	// 匹配链接地址
	urls := xurls.Strict.FindAllString(content, -1)
	if len(urls) == 0 {
		return
	}

	args := map[string]string{
		"url": urls[0],
	}

	reqUrl := d.splitReqUrl("vip-url", args)
	if reqUrl == "" {
		return
	}

	response, err := utils.HttpRequest.Get(reqUrl)
	if err != nil {
		return
	}

	var res resp.DdxVipUrlResponse
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}

	return strings.Replace(content, urls[0], res.Data.Url, -1)

}

// 请求地址
func (d *Ddx) splitReqUrl(m string, args map[string]string) (reqUrl string) {
	method := d.getMethod(m)
	if method == "" {
		return
	}

	u := url.URL{
		Scheme: "http",
		Host:   d.Gateway,
		Path:   method,
	}

	values := url.Values{}

	values.Add("apikey", d.ApiKey)
	for i, arg := range args {
		values.Add(i, arg)
	}

	u.RawQuery = values.Encode()
	return u.String()
}

// 获取请求方法
func (d *Ddx) getMethod(key string) string {
	return viper.GetString("tbk.ddx.method." + key)
}
