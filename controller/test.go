/**
 * @Author: YMBoom
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:12
 */
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

type TestController struct {
}

func (tc TestController) Test(c *gin.Context) {

	u := "http://v2.api.haodanku.com/pdd_goods_search?"
	keyword := "https://u.jd.com/t21rj4z"
	param := url.Values{}
	param.Add("apikey", "jxhw")
	param.Add("keyword", keyword)
	fmt.Println(u + param.Encode())
}
