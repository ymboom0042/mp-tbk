/**
 * @Author: YMBoom
 * @Description:
 * @File:  record
 * @Version: 1.0.0
 * @Date: 2022/02/10 15:26
 */
package model

import (
	"fmt"
	"mp-17208-top/pkg/global"
	"time"
)

type Record struct {
	FromUsername string
	ReceiveContent string
	ReplyContent string
	ApiPlatform uint
	MallPlatform uint
	IsReplace bool
	CreatedAt int64
}

// 获取记录
func GetRecordByReceiveContent(receiveContent string,field string) *Record {
	var r Record
	if err := global.DB.Where("receive_content = ?",receiveContent).
		Select(field).Order("created_at desc").First(&r).Error; err != nil {
			return nil
	}
	return &r
}


// 创建记录
func CreateRecord(fromUsername,ReceiveContent,ReplyContent string,apiPlatform,mallPlatform uint) bool {
	record := &Record{
		FromUsername:   fromUsername,
		ReceiveContent: ReceiveContent,
		ReplyContent:   ReplyContent,
		ApiPlatform:    apiPlatform,
		MallPlatform:   mallPlatform,
		CreatedAt:      time.Now().Unix(),
	}

	if ReceiveContent != ReplyContent {
		record.IsReplace = true
	}

	return record.create()
}

func (r *Record) create() bool {
	if err := global.DB.Create(r).Error; err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
