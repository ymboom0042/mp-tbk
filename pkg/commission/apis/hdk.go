/**
 * @Author: YMBoom
 * @Description:
 * @File:  hdk
 * @Version: 1.0.0
 * @Date: 2022/02/10 10:35
 */
package apis

import (
	"encoding/json"
	"fmt"
	"github.com/mvdan/xurls"
	"github.com/spf13/viper"
	"mp-17208-top/pkg/resp"
	"mp-17208-top/pkg/utils"
	"strings"
)

type Hdk struct {
	Gateway, ApiKey, JdUnionId, VipPid string
}

// 获取京东商品链接
func (h *Hdk) JdItemLink(content string) (tkl string) {
	if content == "" {
		return
	}

	// 匹配链接地址
	urls := xurls.Strict.FindAllString(content, -1)
	if len(urls) == 0 {
		return
	}

	args := map[string]string{
		"apikey":      h.ApiKey,
		"material_id": urls[0],
		"union_id":    h.JdUnionId,
	}

	reqUrl := h.splitReqUrl("jd-link")
	if reqUrl == "" {
		return
	}

	response, err := utils.HttpRequest.Post(reqUrl, args)
	if err != nil {
		return
	}

	var res resp.HdkGetJdItemLinkResponse
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}

	return strings.Replace(content, urls[0], res.Data.ShortUrl, -1)

}

// 获取京东商品链接
func (h *Hdk) VipItemLink(content string) (tkl string) {
	if content == "" {
		return
	}

	// 匹配链接地址
	urls := xurls.Strict.FindAllString(content, -1)
	if len(urls) == 0 {
		return
	}

	args := map[string]string{
		"apikey":  h.ApiKey,
		"goodsid": urls[0],
		"pid":     h.VipPid,
	}

	reqUrl := h.splitReqUrl("vip-link")
	if reqUrl == "" {
		return
	}

	response, err := utils.HttpRequest.Post(reqUrl, args)
	if err != nil {
		return
	}
	fmt.Println(string(response))

	var res resp.HdkGetVipItemLinkResponse
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}

	return strings.Replace(content, urls[0], res.Data.Url, -1)
}

func (h *Hdk) splitReqUrl(m string) (url string) {
	method := h.getMethod(m)
	if method == "" {
		return
	}
	return fmt.Sprintf("%s/%s", h.Gateway, method)
}

// 获取请求方法
func (h *Hdk) getMethod(key string) string {
	return viper.GetString("tbk.hdk.method." + key)
}
