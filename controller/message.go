/**
 * @Author: YMBoom
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2022/01/13 14:51
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ymboom0042/mp-tbk/pkg/service/api"
	"github.com/ymboom0042/mp-tbk/pkg/utils"
)

type MessageController struct {
}

// 接收消息 自动回复消息
func (mc MessageController) Receive(c *gin.Context) {
	var message api.ReceiveMessage
	if err := c.ShouldBindXML(&message); err != nil {
		utils.Println("消息接受失败 err = ", err)
		return
	}

	message.ReplyMsg(c)
}
