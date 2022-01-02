/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 09:09:36
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 09:09:36
 */

package models

import (
	"xcms/consts"
)

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code consts.JsonResultCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

type ListJsonResult struct {
	Code  consts.JsonResultCode `json:"code"`
	Msg   string                `json:"msg"`
	Count int64                 `json:"count"`
	Data  interface{}           `json:"data"`
}
