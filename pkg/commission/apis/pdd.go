/**
 * @Author: YMBoom
 * @Description:
 * @File:  pdd
 * @Version: 1.0.0
 * @Date: 2022/04/22 16:36
 */
package apis

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/mvdan/xurls"
	"io"
	"mp-17208-top/pkg/resp"
	"mp-17208-top/pkg/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Pdd struct {
	Gateway      string
	ClientId     string //必填	POP分配给应用的client_id
	ClientSecret string //必填	POP分配给应用的client_secret
	Pid          string
}

// PddItemLink 拼多多商品转链
func (p *Pdd) PddItemLink(content string) (tkl string) {
	if content == "" {
		return
	}

	// 匹配链接地址
	urls := xurls.Strict.FindAllString(content, -1)
	if len(urls) == 0 {
		return
	}

	args := p.generateReq("pdd.ddk.goods.zs.unit.url.gen", urls[0])
	response, err := utils.HttpRequest.Post(p.Gateway, args)
	if err != nil {
		return
	}

	var res resp.PddGetItemLinkResp
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}

	return strings.Replace(content, urls[0], res.GoodsZsUnitGenerateResponse.MobileShortUrl, -1)
}

// generateReq 生成请求参数
func (p *Pdd) generateReq(method string, sourceUrl string) (params map[string]string) {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	params = map[string]string{
		"source_url": sourceUrl,
		"pid":        p.Pid,
		"type":       method,
		"timestamp":  strconv.FormatInt(loc.Unix(), 10),
		"client_id":  p.ClientId,
		"data_type":  "JSON",
		"version":    "v1",
	}

	params["sign"] = sign(params, p.ClientSecret)
	return
}

// 签名
func sign(args map[string]string, clientSecret string) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range args {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接参数
	query := bytes.NewBufferString(clientSecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(args[k])
	}
	query.WriteString(clientSecret)
	// MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
