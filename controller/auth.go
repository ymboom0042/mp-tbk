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
	"log"
	"mp-17208-top/pkg/utils"
)

type AuthController struct {

}

func (a AuthController) Auth(c *gin.Context)  {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	token := utils.GetConfString("wx.token")
	if ok := utils.CheckSignature(signature, timestamp, nonce, token); !ok {
		log.Println("微信公众号接入校验失败!")
		return
	}

	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr)

}
