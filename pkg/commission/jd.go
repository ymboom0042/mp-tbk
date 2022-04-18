/**
 * @Author: YMBoom
 * @Description:
 * @File:  jd
 * @Version: 1.0.0
 * @Date: 2022/02/10 10:38
 */
package commission

import (
	"github.com/ymboom0042/mp-tbk/pkg/global"
)

type Jd struct {
}

func NewJd() *Jd {
	return &Jd{}
}

func (j *Jd) GetItemLink(content string) (replaceContent string, apiPlatform uint) {
	return global.Hdk.JdItemLink(content), global.ApiPlatformHdk
}
