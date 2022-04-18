/**
 * @Author: YMBoom
 * @Description:
 * @File:  ddx
 * @Version: 1.0.0
 * @Date: 2022/04/18 15:19
 */
package resp

type DdxResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 订单侠唯品会转链返回
type DdxVipUrlResponse struct {
	Data DdxVipUrlResponseData `json:"data"`
}

type DdxVipUrlResponseData struct {
	Url string `json:"url"`
}
