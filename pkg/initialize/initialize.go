/**
 * @Author: YMBoom
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:37
 */
package initialize

import (
	"github.com/spf13/viper"
	"github.com/ymboom0042/mp-tbk/pkg/commission/apis"
	"github.com/ymboom0042/mp-tbk/pkg/global"
	"github.com/ymboom0042/mp-tbk/pkg/utils"
)

func Initialize() {
	loadConfig()

	// 折淘客
	global.Ztk = ztk()
	// 好单库
	global.Hdk = hdk()
	// 订单侠
	global.Ddx = ddx()

	httpRequest()
}

func loadConfig() {
	// 设置配置文件路径
	viper.SetConfigFile(global.ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件加载失败 :" + err.Error())
	}

	utils.Println("配置文件加载成功...")
}

func ztk() *apis.Ztk {
	conf := viper.GetStringMapString("tbk.ztk")
	if conf == nil {
		panic("ztk配置不能为空...")
	}

	return &apis.Ztk{
		Gateway: conf["gateway"],
		AppKey:  conf["app-key"],
		Sid:     conf["sid"],
		Pid:     conf["pid"],
	}
}

func hdk() *apis.Hdk {
	conf := viper.GetStringMapString("tbk.hdk")
	if conf == nil {
		panic("好单库配置不能为空...")
	}

	return &apis.Hdk{
		Gateway:   conf["gateway"],
		ApiKey:    conf["api-key"],
		JdUnionId: conf["jd-union-id"],
		VipPid:    conf["vip-pid"],
	}
}

func ddx() *apis.Ddx {
	conf := viper.GetStringMapString("tbk.ddx")
	if conf == nil {
		panic("订单侠配置不能为空...")
	}

	return &apis.Ddx{
		Gateway: conf["gateway"],
		ApiKey:  conf["api-key"],
	}
}

func httpRequest() {
	utils.NewHttpReq(5)
}
