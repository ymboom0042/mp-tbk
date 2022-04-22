/**
 * @Author: YMBoom
 * @Description:
 * @File:  global
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:38
 */
package global

import (
	"github.com/jinzhu/gorm"
	"mp-17208-top/pkg/commission/apis"
)

var (
	DB  *gorm.DB
	Ztk *apis.Ztk
	Hdk *apis.Hdk
	Ddx *apis.Ddx
	Pdd *apis.Pdd
)
