/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 09:03:35
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 09:03:35
 */
package consts

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)
