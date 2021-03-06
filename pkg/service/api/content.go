/**
 * @Author: YMBoom
 * @Description:
 * @File:  content
 * @Version: 1.0.0
 * @Date: 2022/02/10 14:42
 */
package api

import (
	"mp-17208-top/pkg/commission"
	"mp-17208-top/pkg/global"
	"strings"
)

type Content string

// 替换成自己的内容
func (con Content) ReplaceContent() (content string, apiPlatform, mallPlatform uint) {
	// 匹配渠道
	iCommission, mallPlatform := con.matchPlatform()
	if iCommission == nil {
		return con.toString(), 0, mallPlatform
	}

	content, apiPlatform = iCommission.GetItemLink(con.toString())
	return
}

// 匹配渠道
func (con Content) matchPlatform() (iCommission commission.ICommission, mallPlatform uint) {
	switch {
	//case con.matchTb():
	//	iCommission = commission.NewTb()
	//	mallPlatform = global.MallPlatformTb
	case con.matchJd():
		iCommission = commission.NewJd()
		mallPlatform = global.MallPlatformJd
	case con.matchVip():
		iCommission = commission.NewVip()
		mallPlatform = global.MallPlatformVip
	case con.matchPdd():
		iCommission = commission.NewPdd()
		mallPlatform = global.MallPlatformPdd
	default:
		iCommission = commission.NewTb()
		mallPlatform = global.MallPlatformTb
		return
	}
	return
}

// 匹配淘宝
func (con Content) matchTb() (ok bool) {
	return strings.Contains(con.toString(), "tb.cn")
}

// 匹配京东
func (con Content) matchJd() (ok bool) {
	return strings.Contains(con.toString(), "jd.com")
}

// 匹配唯品会
func (con Content) matchVip() (ok bool) {
	return strings.Contains(con.toString(), "vip.com") ||
		strings.Contains(con.toString(), "vipglobal.hk")
}

// 匹配拼多多
func (con Content) matchPdd() (ok bool) {
	return strings.Contains(con.toString(), "yangkeduo.com") ||
		strings.Contains(con.toString(), "pingduoduo.com")
}

func (con Content) toString() string {
	return string(con)
}
