/**
 * @Author: YMBoom
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:25
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ymboom0042/mp-tbk/pkg/utils"
)

type AuthController struct {
}

func (a AuthController) Auth(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	token := utils.GetConfString("wx.token")
	if ok := utils.CheckSignature(signature, timestamp, nonce, token); !ok {
		return
	}

	_, _ = c.Writer.WriteString(echostr)

}
