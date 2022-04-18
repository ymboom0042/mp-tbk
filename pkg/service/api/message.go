/**
 * @Author: YMBoom
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2022/01/13 14:52
 */
package api

import (
	"encoding/xml"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ymboom0042/mp-tbk/pkg/global"
	"github.com/ymboom0042/mp-tbk/pkg/utils"
	"time"
)

// ReceiveMessage 微信文本消息结构体
type ReceiveMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      Content
	Event        string
	MsgId        int64
}

// 微信回复文本消息结构体
type ReplyTextMessage struct {
	replyMessage
	Content string
}

// 微信回复图片消息结构体
type ReplyImgMessage struct {
	replyMessage
	Image replyImage
}

// 图片
type replyImage struct {
	MediaId string
}

// 返回消息
type replyMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

func (receive *ReceiveMessage) ReplyMsg(c *gin.Context) {
	switch receive.MsgType {
	case global.MsgTypeText:
		receive.text(c)

	case global.MsgTypeEvent:
		if receive.Event == global.MsgEventSubscribe {
			utils.Println("感谢订阅")
			receive.img(c)
		}
	}
}

// 文本消息
func (receive *ReceiveMessage) text(c *gin.Context) {

	// 替换内容
	content, _, _ := receive.Content.ReplaceContent()

	// 返回文本消息
	msg, err := replyTextMsg(receive.FromUserName, receive.ToUserName, content)
	if err != nil {
		return
	}
	_, _ = c.Writer.Write(msg)
}

func (receive *ReceiveMessage) img(c *gin.Context) {
	// 返回文本消息
	msg, err := replyImgMsg(receive.FromUserName, receive.ToUserName, "JNXBA_9MecHTCYS-KW4Dkbb3ZR9Y3eLJkMc5J4_XzaU")
	if err != nil {
		return
	}
	_, _ = c.Writer.Write(msg)
}

// 返回文本消息
func replyTextMsg(toUserName, fromUserName, content string) ([]byte, error) {
	send := ReplyTextMessage{
		replyMessage: replyMessage{
			ToUserName:   toUserName,
			FromUserName: fromUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      global.MsgTypeText,
			XMLName:      xml.Name{},
		},
		Content: content,
	}

	b, err := xml.Marshal(&send)
	if err != nil {
		utils.Println("回复消息xml错误 err= ", err)
		return []byte{}, errors.New("ReplyTextMsgToWx -> xml.Marshal fail")
	}

	return b, nil
}

// 返回图片消息
func replyImgMsg(toUserName, fromUserName, mediaId string) ([]byte, error) {
	send := ReplyImgMessage{
		replyMessage: replyMessage{
			ToUserName:   toUserName,
			FromUserName: fromUserName,
			CreateTime:   time.Now().Unix(),
			MsgType: "image	",
			XMLName: xml.Name{},
		},
		Image: replyImage{MediaId: mediaId},
	}

	utils.Println("send = ", send)
	b, err := xml.Marshal(&send)
	if err != nil {
		utils.Println("回复消息xml错误 err= ", err)
		return []byte{}, errors.New("ReplyTextMsgToWx -> xml.Marshal fail")
	}

	return b, nil
}
