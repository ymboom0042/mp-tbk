/**
 * @Author: YMBoom
 * @Description:
 * @File:  pdd
 * @Version: 1.0.0
 * @Date: 2022/04/22 16:34
 */
package commission

import "mp-17208-top/pkg/global"

type Pdd struct {
}

func NewPdd() *Pdd {
	return &Pdd{}
}

func (p *Pdd) GetItemLink(content string) (replaceContent string, apiPlatform uint) {
	return global.Pdd.PddItemLink(content), global.ApiPlatformPdd
}
