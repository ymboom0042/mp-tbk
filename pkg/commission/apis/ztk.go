/**
 * @Author: YMBoom
 * @Description:
 * @File:  ztk
 * @Version: 1.0.0
 * @Date: 2022/02/10 10:05
 */
package apis

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/ymboom0042/mp-tbk/pkg/resp"
	"github.com/ymboom0042/mp-tbk/pkg/utils"
)

type Ztk struct {
	Gateway, AppKey, Sid, Pid string
}

// 淘宝链接
func (z *Ztk) TbItemLink(content string) (tkl string) {
	if content == "" {
		return
	}

	args := map[string]string{
		"tkl":     content,
		"signurl": "4",
	}

	reqUrl := z.splitReqUrl("get-tkl", args)
	if reqUrl == "" {
		return
	}

	response, err := utils.HttpRequest.Get(reqUrl)
	if err != nil {
		return
	}

	var res resp.ZtkGetTklResponse
	err = json.Unmarshal(response, &res)
	if err != nil {
		fmt.Println("json err = ", err)
		return
	}
	tkl = res.TbkPrivilegeGetResponse.Result.Data.Tkl
	if tkl == "" {
		return
	}
	return z.splitRetContent(tkl, res.TbkPrivilegeGetResponse.Result.Data.Title)

}

// 返回数据
func (z *Ztk) splitRetContent(tkl, title string) (str string) {
	str += title + "\n"
	str += "fu 制" + "\n"
	str += tkl + "打開 tao宝"
	return
}

// 请求地址
func (z *Ztk) splitReqUrl(m string, args map[string]string) (url string) {
	method := z.getMethod(m)
	if method == "" {
		return
	}

	return fmt.Sprintf("%s/%s.ashx?%s", z.Gateway, method, z.argsEncode(args))
}

// 获取请求方法
func (z *Ztk) getMethod(key string) string {
	return viper.GetString("tbk.ztk.method." + key)
}

// 参数
func (z *Ztk) argsEncode(args map[string]string) string {
	if args == nil {
		return ""
	}

	args["appkey"] = z.AppKey
	args["sid"] = z.Sid
	args["pid"] = z.Pid

	return utils.ArgsEncode(args).Encode()
}
