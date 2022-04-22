/**
 * @Author: YMBoom
 * @Description:
 * @File:  jd
 * @Version: 1.0.0
 * @Date: 2022/02/10 10:38
 */
package commission

import (
	"mp-17208-top/pkg/global"
)

type Vip struct {
}

func NewVip() *Vip {
	return &Vip{}
}

func (v *Vip) GetItemLink(content string) (replaceContent string, apiPlatform uint) {
	return global.Ddx.VipItemLink(content), global.ApiPlatformHdk
}
