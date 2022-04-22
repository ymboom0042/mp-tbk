/**
 * @Author: YMBoom
 * @Description:
 * @File:  pdd
 * @Version: 1.0.0
 * @Date: 2022/04/22 17:03
 */
package resp

type PddGetItemLinkResp struct {
	GoodsZsUnitGenerateResponse GoodsZsUnitGenerateResponse `json:"goods_zs_unit_generate_response"`
}

type GoodsZsUnitGenerateResponse struct {
	MobileShortUrl string `json:"mobile_short_url"`
}
