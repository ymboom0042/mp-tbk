/**
 * @Author: YMBoom
 * @Description:
 * @File:  commission
 * @Version: 1.0.0
 * @Date: 2022/02/10 10:04
 */
package commission

type ICommission interface {
	GetItemLink(content string) (string, uint)
}
