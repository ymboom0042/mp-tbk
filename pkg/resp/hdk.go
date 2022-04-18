/**
 * @Author: YMBoom
 * @Description:
 * @File:  hdk
 * @Version: 1.0.0
 * @Date: 2022/02/10 13:16
 */
package resp

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 京东转链返回数据
type HdkGetJdItemLinkResponse struct {
	Response
	Data HdkGetJdItemLinkData `json:"data"`
}

type HdkGetJdItemLinkData struct {
	ShortUrl string `json:"short_url"`
}

// 唯品会转链返回数据
type HdkGetVipItemLinkResponse struct {
	Response
	Data HdkGetVipItemLinkData `json:"data"`
}

type HdkGetVipItemLinkData struct {
	Url string `json:"url"`
}
