/**
 * @Author: YMBoom
 * @Description:
 * @File:  const
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:36
 */
package global

// 配置文件路径
const ConfigPath = "config/config.yaml"

const (
	// 文本消息
	MsgTypeText = "text"
	// 事件
	MsgTypeEvent = "event"
	// 事件-订阅
	MsgEventSubscribe = "subscribe"
	// 事件-(取消订阅
	MsgEventUnSubscribe = "unsubscribe"

	MallPlatformTb  = 1
	MallPlatformJd  = 2
	MallPlatformVip = 3
	MallPlatformPdd = 4

	ApiPlatformZtk = 1
	ApiPlatformHdk = 2
	ApiPlatformDdx = 3
	ApiPlatformPdd = 4
)
