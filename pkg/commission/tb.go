/**
 * @Author: YMBoom
 * @Description:
 * @File:  tb
 * @Version: 1.0.0
 * @Date: 2022/02/10 10:28
 */
package commission

import (
	"github.com/ymboom0042/mp-tbk/pkg/global"
)

type Tb struct {
}

func NewTb() *Tb {
	return &Tb{}
}

// 获取商品链接
func (t *Tb) GetItemLink(content string) (replyContent string, apiPlatform uint) {
	return global.Ztk.TbItemLink(content), global.ApiPlatformZtk
}
