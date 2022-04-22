/**
 * @Author: YMBoom
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:37
 */
package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"mp-17208-top/pkg/commission/apis"
	"mp-17208-top/pkg/global"
	"mp-17208-top/pkg/utils"
	"time"
)

var pddPid = "9742959_237808983"

func Initialize() {
	loadConfig()

	// 数据库
	global.DB = database()

	// 折淘客
	global.Ztk = ztk()
	// 好单库
	global.Hdk = hdk()
	// 订单侠
	global.Ddx = ddx()
	// 拼多多
	global.Pdd = pdd()

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

func database() *gorm.DB {
	mysql := viper.GetStringMap("mysql")
	args := fmt.Sprintf("%s:%s@/%s?%s",
		mysql["username"],
		mysql["password"],
		mysql["dbname"],
		mysql["config"],
	)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err.Error())
	}

	//默认不加复数
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	// sql语句打印
	if mysql["log-mode"] == true {
		db.LogMode(true)
	}

	utils.Println("数据库链接成功...")
	return db
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

func pdd() *apis.Pdd {
	conf := viper.GetStringMapString("tbk.pdd")
	if conf == nil {
		panic("拼多多配置不能为空...")
	}
	fmt.Println(viper.GetString("tbk.pdd.pid"))

	return &apis.Pdd{
		Gateway:      conf["gateway"],
		ClientId:     conf["client-id"],
		ClientSecret: conf["client-secret"],
		Pid:          pddPid,
	}
}

func httpRequest() {
	utils.NewHttpReq(5)
}
