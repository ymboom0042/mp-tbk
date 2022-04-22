/**
 * @Author: YMBoom
 * @Description:
 * @File:  ztk
 * @Version: 1.0.0
 * @Date: 2022/02/10 13:16
 */
package resp

type ZtkGetTklResponse struct {
	TbkPrivilegeGetResponse ZtkTbkPrivilegeGetResponse `json:"tbk_privilege_get_response"`
}

type ZtkTbkPrivilegeGetResponse struct {
	Result ZtkResult `json:"result"`
	RequestId string `json:"request_id"`
}

type ZtkResult struct {
	Data ZtkData `json:"data"`
}

type ZtkData struct {
	Title string `json:"title"`
	Tkl string `json:"tkl"`
}
